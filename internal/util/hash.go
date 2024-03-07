package util

import (
	"crypto/sha256"
	"math/rand"
)

var (
	allowedSymbols = []rune(
		"abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789",
	)
	tokenSize = 6
)

func GetStringIntHash(s string) int64 {
	hash := sha256.New()
	hash.Write([]byte(s))

	hashBytes := hash.Sum(nil)

	var result int64 = 0

	for i := 0; i < 8; i++ {
		result <<= 8
		result += int64(hashBytes[i])
	}

	return result
}

func GetToken(s string) string {
	symbols := make([]rune, tokenSize)

	random := rand.New(rand.NewSource(GetStringIntHash(s)))

	for i := 0; i < tokenSize; i++ {
		symbols[i] = allowedSymbols[random.Intn(len(allowedSymbols))]
	}

	return string(symbols)
}
