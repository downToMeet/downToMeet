package impl_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.timothygu.me/downtomeet/server/impl"
)

func TestImplementation_NewOAuthState(t *testing.T) {
	s := testImpl.NewOAuthState()
	assert.Len(t, s.State, StateNonceLen)
	assert.False(t, s.ExpiresAt.After(time.Now().Add(StateExpiry)))
}

func TestOAuthState_Validate(t *testing.T) {
	const (
		state    = "123"
		notstate = "456"
	)
	s := OAuthState{
		State:     state,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	assert.True(t, s.Validate(state))
	assert.False(t, s.Validate(notstate))
}

func TestOAuthState_Validate_Expired(t *testing.T) {
	const state = "123"
	s := OAuthState{
		State:     state,
		ExpiresAt: time.Now().Add(-24 * time.Hour),
	}
	assert.False(t, s.Validate(state))
}
