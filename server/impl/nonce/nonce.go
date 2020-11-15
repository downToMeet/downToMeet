package nonce

import (
	"math/rand"
	"sync"
)

// Generator creates randomly generated ASCII nonces. The set of possible
// characters in generated nonces consists of alphanumerics plus _ -.
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

// NewState creates a new nonce and remembers it for later retrieval.
func (g *Generator) NewState(stateLen int) string {
	g.mu.Lock()
	defer g.mu.Unlock()

	return RandomBase64(g.rand, stateLen)
}
