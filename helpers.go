package main

import (
	"math/rand"
)

func generateCode() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	str := ""

	for i := 0; i < 6; i++ {
		str += string(charset[rand.Intn(len(charset))])

	}
	return str

}
