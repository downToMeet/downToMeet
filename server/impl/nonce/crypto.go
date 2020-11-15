package nonce

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	mathrand "math/rand"
	"sync"
)

const cryptoRandBufferSize = 128

// cryptoRandSource is an implementation of rand.Source64 backed by crypto/rand.
type cryptoRandSource struct {
	r   *bufio.Reader
	buf []byte
}

var (
	_ mathrand.Source   = cryptoRandSource{}
	_ mathrand.Source64 = cryptoRandSource{}
)

var testCryptoRandAvailability sync.Once

func NewCryptoRandSource() mathrand.Source64 {
	testCryptoRandAvailability.Do(func() {
		buf := make([]byte, 1)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			panic(fmt.Errorf("nonce: crypto/rand is unavailable: %w", err))
		}
	})

	return cryptoRandSource{
		r:   bufio.NewReaderSize(rand.Reader, cryptoRandBufferSize),
		buf: make([]byte, 64/8),
	}
}

func (s cryptoRandSource) Seed(_ int64) {
	// crypto/rand doesn't support seeding.
}

func (s cryptoRandSource) Int63() int64 {
	n := s.Uint64()
	return int64(n &^ (1 << 63))
}

func (s cryptoRandSource) Uint64() uint64 {
	_, err := io.ReadFull(s.r, s.buf)
	if err != nil {
		panic(fmt.Errorf("nonce: failed to read from crypto/rand: %w", err))
	}
	return binary.LittleEndian.Uint64(s.buf)
}
