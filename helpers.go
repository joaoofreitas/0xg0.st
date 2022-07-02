package main

import "math/rand"

// Generate a UUID
func GenerateUUID() string {
	var symbols = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	var uuid string
	for i := 0; i < 12; i++ {
		uuid += string(symbols[rand.Intn(len(symbols)-1)])
	}

	return uuid
}
