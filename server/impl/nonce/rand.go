package nonce

import (
	"encoding/base32"
	"math/rand"
	"strings"
)

var base32Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)

func RandomAlphanumerics(r *rand.Rand, n int) string {
	// Extra bytes to allocate for base32 overhead.
	const (
		inputOverhead  = 1
		outputOverhead = 2
	)

	var b strings.Builder
	b.Grow(n + outputOverhead)
	enc := base32.NewEncoder(base32Encoding, &b)
	randBytes := make([]byte, base32Encoding.DecodedLen(n)+inputOverhead)

	r.Read(randBytes)
	_, _ = enc.Write(randBytes)
	enc.Close()
	return b.String()[:n]
}
