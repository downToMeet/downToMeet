// Package nonce implements a simple ASCII nonce generator, as well as some
// randomization utilities.
package nonce

import (
	"math/rand"
	"sync"
)

// Generator creates randomly generated ASCII nonces. The set of possible
// characters in generated nonces is the same as RandomBase64. A Generator
// instance is safe for concurrent use.
type Generator struct {
	rand *rand.Rand
	mu   sync.Mutex
}

// NewGenerator creates a new Generator from the given rand.Source.
func NewGenerator(src rand.Source) *Generator {
	return &Generator{
		rand: rand.New(src),
	}
}

// NewBase64 creates a new nonce with the given length.
func (g *Generator) NewBase64(stateLen int) string {
	g.mu.Lock()
	defer g.mu.Unlock()

	return RandomBase64(g.rand, stateLen)
}
