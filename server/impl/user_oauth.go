package impl

import "time"

const (
	stateNonceLen = 30
	stateExpiry   = 30 * time.Minute // login must be complete within 30 min
)

// OAuthState represents the current OAuth 2.0 login session. In an OAuth login
// flow, this structure would be created upon user commencing a login, with the
// State field used as the "state" URL parameter and the entire structure saved
// to the user's cookie. Later, when the user redirects back to the server, the
// server would validate the validity of the state by comparing the received
// "state" parameter against the State field.
type OAuthState struct {
	State     string // a nonce used as the OAuth "state" parameter
	ExpiresAt time.Time
}

// Validate returns true if s is not expired, and the given string matches
// s.State.
func (s OAuthState) Validate(state string) bool {
	return s.State == state && s.ExpiresAt.After(time.Now())
}

// NewOAuthState returns a fresh OAuthState. The State field is set to a
// randomly generated 30-character ASCII string. The returned state is set to
// expire after 30 minutes.
func (i *Implementation) NewOAuthState() OAuthState {
	return OAuthState{
		State:     i.nonceGen.NewBase64(stateNonceLen),
		ExpiresAt: time.Now().UTC().Add(stateExpiry),
	}
}
