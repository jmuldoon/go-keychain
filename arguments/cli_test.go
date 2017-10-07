package arguments

import "testing"

const targetTestVersion = 0

func TestVersionValidation(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v\n", testVersion, targetTestVersion)
	}
}
