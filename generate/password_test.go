package generate

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

const targetTestVersion = 0

func TestVersionValidation(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v\n", testVersion, targetTestVersion)
	}
}

func TestRandStringBytesMaskLength(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	for _, test := range testStringLengthCases {
		observed := RandStringBytesMask(src, test.Tested.Value.(int))
		t.Logf("Running test for `%s`\n", test.Description)
		if len(observed) != test.Expected.Value {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestRandStringBytesMask", test.Tested.Value,
				test.Description, observed, test.Expected.Value)
		}
	}
}

func TestRandStringBytesMaskType(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	for _, test := range testStringLengthCases {
		observed := RandStringBytesMask(src, test.Tested.Value.(int))
		t.Logf("Running test for `%s`\n", test.Description)
		typecheck := reflect.TypeOf(observed)
		if typecheck.Kind() != reflect.String {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestRandStringBytesMask", test.Tested.Value,
				test.Description, observed, test.Expected.Value)
		}
	}
}

// TODO: figure out why benchmarks are not working
func benchmarkRandStringBytesMask(n int, b *testing.B) {
	// Create the new source for random generation
	src := rand.NewSource(time.Now().UnixNano())
	// Run the RandStringBytesMask function b.N times
	for n := 0; n < b.N; n++ {
		RandStringBytesMask(src, n)
	}
}

func BenchmarkRandStringBytesMask2(b *testing.B)  { benchmarkRandStringBytesMask(2, b) }
func BenchmarkRandStringBytesMask4(b *testing.B)  { benchmarkRandStringBytesMask(4, b) }
func BenchmarkRandStringBytesMask8(b *testing.B)  { benchmarkRandStringBytesMask(8, b) }
func BenchmarkRandStringBytesMask16(b *testing.B) { benchmarkRandStringBytesMask(16, b) }
func BenchmarkRandStringBytesMask32(b *testing.B) { benchmarkRandStringBytesMask(32, b) }
func BenchmarkRandStringBytesMask64(b *testing.B) { benchmarkRandStringBytesMask(64, b) }
