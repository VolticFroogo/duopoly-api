package helper

import (
	"math/rand"
	"time"
)

const (
	letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	dieFaces = 6
)

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func RandomDice() [2]int {
	return [2]int{
		rand.Intn(dieFaces) + 1,
		rand.Intn(dieFaces) + 1,
	}
}
