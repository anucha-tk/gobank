package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max
func RandomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random string by n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Random Owner string by 6 char
func RandomOwner() string {
	return RandomString(6)
}

// Random Money between 0-1000
func RandomMoney() int64 {
	return RandomInit(0, 1000)
}

// Random Currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
