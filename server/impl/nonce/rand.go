package nonce

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

var base64Encoding = base64.RawURLEncoding

// RandomBase64 returns a string of length n, consisting of random characters
// drawn from the provided rand.Rand. The output string will solely consist of
// base64url characters; that is, all alphanumerics with - and _ added.
func RandomBase64(r *rand.Rand, n int) string {
	// Extra bytes to allocate for base64 overhead.
	const (
		inputOverhead  = 1
		outputOverhead = 1
	)

	var b strings.Builder
	b.Grow(n + outputOverhead)
	enc := base64.NewEncoder(base64Encoding, &b)
	randBytes := make([]byte, base64Encoding.DecodedLen(n)+inputOverhead)

	r.Read(randBytes)
	_, _ = enc.Write(randBytes)
	enc.Close()
	return b.String()[:n]
}
