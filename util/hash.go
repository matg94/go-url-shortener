package util

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/matg94/go-url-shortener/errorhandling"
)

func HashString(original string, hashLength int) string {
	hasher := sha1.New()
	_, err := hasher.Write([]byte(original))
	if err != nil {
		errorhandling.HandleError(err, "Hashing", original)
		return ""
	}
	sha := hex.EncodeToString(hasher.Sum(nil)[:hashLength])
	return sha
}
