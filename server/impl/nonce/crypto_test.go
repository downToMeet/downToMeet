package nonce_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "go.timothygu.me/downtomeet/server/impl/nonce"
)

func TestNewCryptoRandSource(t *testing.T) {
	const iterations = 100

	s := NewCryptoRandSource()
	s.Uint64()
	for i := 0; i < iterations; i++ {
		require.GreaterOrEqual(t, s.Int63(), int64(0))
	}
}

func BenchmarkNewCryptoRandSource(b *testing.B) {
	b.ReportAllocs()

	s := NewCryptoRandSource()
	for i := 0; i < b.N; i++ {
		s.Uint64()
	}
}
