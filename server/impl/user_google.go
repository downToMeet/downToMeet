package impl

import (
	"fmt"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"gopkg.in/square/go-jose.v2/jwt"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

var (
	googleConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     endpoints.Google,
		// https://developers.google.com/identity/protocols/oauth2/scopes#google-sign-in
		Scopes: []string{"openid", "profile", "email"},
	}
)

func (i *Implementation) GetUserGoogleAuth(param operations.GetUserGoogleAuthParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()
	session := SessionFromContext(ctx)

	oauthState := i.NewOAuthState()
	session.Values[GoogleState] = oauthState

	config := i.buildOAuthConfig(param.HTTPRequest, googleConfig, new(operations.GetUserGoogleRedirectURL))
	return operations.NewGetUserGoogleAuthSeeOther().
		WithLocation(config.AuthCodeURL(oauthState.State))
}

func (i *Implementation) GetUserGoogleRedirect(param operations.GetUserGoogleRedirectParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()
	logger := log.WithContext(ctx)
	session := SessionFromContext(ctx)

	// If we don't find a cookie, and we haven't had tried the trampoline trick
	// yet, try to do a soft redirect to ourselves. Hopefully, this will make
	// browsers think we are in a same-site context.
	if session.IsNew {
		if swag.StringValue(param.Trampoline) == "" {
			logger.Info("no cookie found; try trampoline")
			u := operations.GetUserGoogleRedirectURL{
				Code:       param.Code,
				State:      param.State,
				Trampoline: swag.String("1"),
			}
			return SoftRedirect{i.buildURL(param.HTTPRequest, &u)}
		}
		logger.Warn("no cookie found, but already tried trampoline")
	}

	redirectToHome := operations.NewGetUserGoogleRedirectSeeOther().
		WithLocation(i.buildURL(param.HTTPRequest, &operations.GetUserMeURL{}))

	// Step 1: Check request state validity to protect against CSRF attacks.
	// See https://auth0.com/docs/protocols/state-parameters.

	storedState := session.Values[GoogleState]
	delete(session.Values, GoogleState)
	if storedState == nil || !storedState.(OAuthState).Validate(param.State) {
		logger.WithFields(log.Fields{"expected_state": storedState, "got_state": param.State}).
			Error("Invalid state when logging into Google")
		return redirectToHome
	}

	// Step 2: Exchange the received authorization code with Google for an
	// access token.

	config := i.buildOAuthConfig(param.HTTPRequest, googleConfig, new(operations.GetUserGoogleRedirectURL))
	token, err := config.Exchange(ctx, param.Code)
	if err != nil {
		logger.WithError(err).Error("Failed to exchange for access token from Google")
		return redirectToHome
	} else if !token.Valid() {
		logger.WithFields(log.Fields{
			"token":      token,
			"token_type": token.TokenType,
			"access":     token.AccessToken,
			"refresh":    token.RefreshToken,
			"expiry":     token.Expiry,
			"id_token":   token.Extra("id_token"),
		}).Error("Google provided invalid access token")
		return redirectToHome
	}

	// We don't have to validate the token here since we just got it from Google
	// over an encrypted channel.
	claims, err := unsafeGoogleJWTClaims(token)
	if err != nil {
		logger.WithError(err).WithFields(log.Fields{
			"token":      token,
			"token_type": token.TokenType,
			"access":     token.AccessToken,
			"refresh":    token.RefreshToken,
			"expiry":     token.Expiry,
			"id_token":   token.Extra("id_token"),
			"claims":     claims,
		}).Error("Google did not provide a valid id_token")
		return redirectToHome
	}

	logger.WithFields(log.Fields{
		"token_type": token.TokenType,
		"access":     token.AccessToken,
		"refresh":    token.RefreshToken,
		"expiry":     token.Expiry,
		"claims":     claims,
	}).Info("Successfully logged into Google")

	// Step 3.1: Match the Google user with an existing user, if possible.

	tx := i.DB().WithContext(ctx)

	// Look for a user with this Google ID.
	var dbUser db.User
	if err := tx.First(&dbUser, "google_id = ?", claims.Subject).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.WithError(err).Error("Unable to lookup user by Google ID")
		return redirectToHome
	} else if err == nil {
		session.Values[UserID] = dbUser.IDString()
		logger.Info("logged in through facebook ID")
		return redirectToHome
	}

	// Look for a user with this email if they have a verified email. If found,
	// add the Google ID to their account.
	if claims.Email != "" && claims.EmailVerified {
		if err := tx.First(&dbUser, "email = ?", claims.Email).Error; err != nil && err != gorm.ErrRecordNotFound {
			logger.WithError(err).Error("Unable to lookup user by email")
			return redirectToHome
		} else if err == nil {
			session.Values[UserID] = dbUser.IDString()

			if dbUser.GoogleID == nil {
				dbUser.GoogleID = swag.String(claims.Subject)
				if err := tx.Model(&dbUser).Update("GoogleID", dbUser.GoogleID).Error; err != nil {
					logger.WithError(err).Warn("Unable to update user's Google ID")
				}
			}

			if dbUser.ProfilePic == nil && claims.PictureURL != "" {
				dbUser.ProfilePic = swag.String(claims.PictureURL)
				if err := tx.Model(&dbUser).Update("ProfilePic", dbUser.ProfilePic).Error; err != nil {
					logger.WithError(err).Warn("Unable to update user's profile picture")
				}
			}

			logger.Info("logged in through email")
			return redirectToHome
		}
	}

	// Step 3.2: Create a new user for this Google user.

	dbUser.Name = claims.Name
	if claims.EmailVerified {
		dbUser.Email = claims.Email
	}
	dbUser.GoogleID = swag.String(claims.Subject)
	if claims.PictureURL != "" {
		dbUser.ProfilePic = swag.String(claims.PictureURL)
	}
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

// https://developers.google.com/identity/protocols/oauth2/openid-connect#an-id-tokens-payload
type googleClaims struct {
	jwt.Claims
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	HostedDomain  string `json:"hd,omitempty"`
	Locale        string `json:"locale,omitempty"` // BCP 47
	Name          string `json:"name,omitempty"`
	Nonce         string `json:"nonce,omitempty"`
	PictureURL    string `json:"picture,omitempty"`
	ProfileURL    string `json:"profile,omitempty"`
}

// unsafeGoogleToken returns the id_token field inside an oauth2.Token.
// It assumes that we got the id_token from a legitimate source, and does not
// attempt to validate the token.
func unsafeGoogleJWTClaims(token *oauth2.Token) (*googleClaims, error) {
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("no id_token provided")
	}

	idTokenParsed, err := jwt.ParseSigned(idToken)
	if err != nil {
		return nil, fmt.Errorf("unable to parse id_token: %w", err)
	}

	var claims googleClaims
	return &claims, idTokenParsed.UnsafeClaimsWithoutVerification(&claims)
}
