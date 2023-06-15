package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func Sha256(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	out := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return out
}
