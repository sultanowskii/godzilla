package util

import (
	"crypto/md5"
	"encoding/hex"
)

func HashString(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
