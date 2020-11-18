package nonce_test

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.timothygu.me/downtomeet/server/impl/nonce"
)

var validNonce = regexp.MustCompile(`^[0-9a-zA-Z_-]*$`)

func TestRandomAlphanumerics(t *testing.T) {
	const seed = 1
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < 20; i++ {
		str := RandomBase64(r, i)
		assert.Len(t, str, i, i)
		assert.Regexp(t, validNonce, str, i)
	}
}
