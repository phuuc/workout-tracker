package helpers

import (
	"math/rand"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number      = []rune("0123456789")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandNumber(n int) string {
	num := make([]rune, n)
	for i := 0; i < n; i++ {
		num[i] = number[rand.Intn(len(number))]
	}
	return string(num)
}
