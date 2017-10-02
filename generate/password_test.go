package generate

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRandStringBytesMask(b *testing.B) {
	// Create the new source for random generation
	src := rand.NewSource(time.Now().UnixNano())
	// Run the RandStringBytesMask function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMask(src, 32)
	}
}

func TestRandStringBytesMask(t *testing.T) {
	// Create the new source for random generation
	src := rand.NewSource(time.Now().UnixNano())
	length := 32
	observed := RandStringBytesMask(src, length)
	if observed == "" || len(observed) != length {
		t.Fatalf("TestRandStringBytesMask failed with\n"+
			"length %d != %d\n"+
			"string %s\n", len(observed), length, observed)
	}
}
