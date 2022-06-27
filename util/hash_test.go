package util

import "testing"

func TestHashString(t *testing.T) {
	testString := "test_string"
	testHashLength := 5
	hash := HashString(testString, testHashLength)

	if len(hash) != (testHashLength * 2) {
		t.Logf("expected hash to be of length %d, but got %d", testHashLength*2, len(hash))
		t.Fail()
	}
}
