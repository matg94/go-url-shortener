package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func HashString(original string, hashLength int) string {
	hasher := sha1.New()
	hasher.Write([]byte(original))
	sha := hex.EncodeToString(hasher.Sum(nil)[:hashLength])
	fmt.Println(sha)
	return sha
}