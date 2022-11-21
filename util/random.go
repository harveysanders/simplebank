package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a string of length length.
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alpha)

	for i := 0; i < length; i++ {
		c := alpha[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random account owner.
func RandomOwner() string {
	return RandomString(6)
}

// RandomAmount generates a random amount of money.
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a currency abbreviation.
func RandomCurrency() string {
	currs := []string{"USD", "EUR", "CAN", "BIR"}
	return currs[rand.Intn(len(currs))]
}
