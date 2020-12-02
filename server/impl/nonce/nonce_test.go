package nonce_test

import (
	"math/rand"
	"sync"
	"testing"

	. "go.timothygu.me/downtomeet/server/impl/nonce"
)

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
				g.NewBase64(stateLen)
			}
		}()
	}
	wg.Wait()
}
