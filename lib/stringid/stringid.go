package stringid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strconv"
	"strings"

	"github.com/nightlegend/hi/lib/random"
)

const shortLen = 12

// TruncateID returns a shorthand version of a string identifier for convenience.
// A collision with other shorthands is very unlikely, but possible.
// In case of a collision a lookup with TruncIndex.Get() will fail, and the caller
// will need to use a longer prefix, or the full-length Id.
func TruncateID(id string) string {
	if i := strings.IndexRune(id, ':'); i >= 0 {
		id = id[i+1:]
	}
	if len(id) > shortLen {
		id = id[:shortLen]
	}
	return id
}

func generateID(crypto bool) string {
	b := make([]byte, 32)
	r := random.Reader
	if crypto {
		r = rand.Reader
	}
	for {
		if _, err := io.ReadFull(r, b); err != nil {
			panic(err) // This shouldn't happen
		}
		id := hex.EncodeToString(b)
		// if we try to parse the truncated for as an int and we don't have
		// an error then the value is all numeric and causes issues when
		// used as a hostname. ref #3869
		if _, err := strconv.ParseInt(TruncateID(id), 10, 64); err == nil {
			continue
		}
		return id
	}
}

// GenerateNonCryptoID generates unique id without using cryptographically
// secure sources of random.
// It helps you to save entropy.
func GenerateNonCryptoID() string {
	return generateID(false)
}
