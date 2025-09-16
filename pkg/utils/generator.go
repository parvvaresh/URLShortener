package utils

import (
	"math/rand"
	"time"
)


const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func init() {
	rand.Seed(time.Now().UnixNano())
}


func GenerateCode(lenght int) string {
	b := make([]byte, lenght)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
