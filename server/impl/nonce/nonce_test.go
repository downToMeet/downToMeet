package nonce_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.timothygu.me/downtomeet/server/impl/nonce"
)

func TestGenerator_NewState_Len(t *testing.T) {
	const seed = 1

	g := NewGenerator(rand.NewSource(seed))
	for i := 0; i < 20; i++ {
		state := g.NewState(i)
		assert.Len(t, state, i, i)
		assert.Regexp(t, alphanumerics, state, i)
		assert.True(t, g.VerifyState(state), state)
		assert.False(t, g.VerifyState(state), state)
		assert.False(t, g.VerifyState(state), state)
	}
}

func TestGenerator_NewState_Concurrency(t *testing.T) {
	const (
		seed        = 1
		concurrency = 100
		repetition  = 100
		stateLen    = 16
	)

	g := NewGenerator(rand.NewSource(seed))
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < repetition; j++ {
				g.NewState(stateLen)
			}
		}()
	}
	wg.Wait()
}
