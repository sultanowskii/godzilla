package util

import (
	"crypto/sha256"
	"math/rand"
)

var (
	allowedSymbols = []rune(
		"abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789",
	)
	dzillaSize = 6
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

func GetDzilla(s string) string {
	symbols := make([]rune, dzillaSize)

	random := rand.New(rand.NewSource(GetStringIntHash(s)))

	for i := 0; i < dzillaSize; i++ {
		symbols[i] = allowedSymbols[random.Intn(len(allowedSymbols))]
	}

	return string(symbols)
}
