package impl

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

var (
	facebookConfig = oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_APP_ID"),
		ClientSecret: os.Getenv("FACEBOOK_APP_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/v9.0/dialog/oauth",
			TokenURL: "https://graph.facebook.com/v9.0/oauth/access_token",
		},
		RedirectURL: "http://localhost:9000/user/facebook/redirect", // TODO: make this generic
		Scopes:      []string{"email"},
	}
)

const (
	stateNonceLen = 30
	stateExpiry   = 30 * time.Minute // login must be complete within 30 min
)

type OAuthState struct {
	State     string
	ExpiresAt time.Time
}

func (s OAuthState) Validate(state string) bool {
	return s.State == state && s.ExpiresAt.After(time.Now())
}

func (i *Implementation) NewOAuthState() OAuthState {
	return OAuthState{
		State:     i.nonceGen.NewState(stateNonceLen),
		ExpiresAt: time.Now().UTC().Add(stateExpiry),
	}
}

func (i *Implementation) GetUserFacebookAuth(param operations.GetUserFacebookAuthParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()
	session := SessionFromContext(ctx)

	oauthState := i.NewOAuthState()
	session.Values[FacebookState] = oauthState

	return operations.NewGetUserFacebookAuthSeeOther().
		WithLocation(facebookConfig.AuthCodeURL(oauthState.State))
}

func (i *Implementation) GetUserFacebookRedirect(param operations.GetUserFacebookRedirectParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()
	logger := log.WithContext(ctx)
	session := SessionFromContext(ctx)
	redirectToHome := operations.NewGetUserFacebookRedirectSeeOther().
		WithLocation(new(operations.GetUserMeURL).String()) // TODO: redirect to actual homepage

	// Step 1: Check request state validity to protect against CSRF attacks.
	// See https://auth0.com/docs/protocols/state-parameters.

	storedState := session.Values[FacebookState]
	delete(session.Values, FacebookState)
	if storedState == nil || !storedState.(OAuthState).Validate(param.State) {
		logger.WithFields(log.Fields{"expected_state": storedState, "got_state": param.State}).
			Error("Invalid state when logging into Facebook")
		return redirectToHome
	}

	// Step 2: Exchange the received authorization code with Facebook for an
	// access token.

	token, err := facebookConfig.Exchange(ctx, param.Code)
	if err != nil {
		logger.WithError(err).Error("Failed to exchange for access token from Facebook")
		return redirectToHome
	} else if !token.Valid() {
		logger.WithField("token", token).Error("Facebook provided invalid access token")
		return redirectToHome
	}

	logger.WithFields(log.Fields{
		"token_type": token.TokenType,
		"access":     token.AccessToken,
		"refresh":    token.RefreshToken,
		"expiry":     token.Expiry,
	}).Info("Successfully logged into Facebook")

	// Step 3: Fetch user information using the access token.

	c := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))
	resp, err := c.Get("https://graph.facebook.com/v9.0/me?fields=id%2Cname%2Cemail&format=json")
	if err != nil {
		logger.WithField("token", token).WithError(err).Error("Unable to get user information from Facebook")
		return redirectToHome
	}
	defer resp.Body.Close()

	jsonDec := json.NewDecoder(resp.Body)
	var info struct{ ID, Name, Email string }
	if err := jsonDec.Decode(&info); err != nil {
		logger.WithField("token", token).WithError(err).Error("Unable to get user information from Facebook")
		return redirectToHome
	}

	logger.Info(info)

	// Step 4.1: Match the Facebook user with an existing user, if possible.

	tx := i.DB().WithContext(ctx)

	// Look for a user with this Facebook ID.
	var dbUser db.User
	if err := tx.First(&dbUser, "facebook_id = ?", info.ID).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.WithError(err).Error("Unable to lookup user by Facebook ID")
		return redirectToHome
	} else if err == nil {
		session.Values[UserID] = dbUser.IDString()
		logger.Info("logged in through facebook ID")
		return redirectToHome
	}

	// Look for a user with this email. If found, add the Facebook ID to their
	// account.
	if err := tx.First(&dbUser, "email = ?", info.Email).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.WithError(err).Error("Unable to lookup user by email")
		return redirectToHome
	} else if err == nil {
		session.Values[UserID] = dbUser.IDString()

		if dbUser.FacebookID == nil {
			dbUser.FacebookID = swag.String(info.ID)
			if err := tx.Model(&dbUser).Update("FacebookID", dbUser.FacebookID).Error; err != nil {
				logger.WithError(err).Warn("Unable to update user's Facebook ID")
			}
		}

		logger.Info("logged in through email")
		return redirectToHome
	}

	// Step 4.2: Create a new user for this Facebook user.

	dbUser.Name = info.Name
	dbUser.FacebookID = swag.String(info.ID)
	dbUser.Email = info.Email
	if err := tx.Create(&dbUser).Error; err != nil {
		logger.WithError(err).Error("Unable to create new user")
		return redirectToHome
	}

	if dbUser.ID == 0 {
		logger.Error("Created user but ID is still empty")
		return redirectToHome
	}

	session.Values[UserID] = dbUser.IDString()

	logger.Info("created new user")
	return redirectToHome
}
