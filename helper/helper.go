package helper

import "math/rand"

const letters = "asdfghjklqwertyuiopzxcvbnm"

func GenerateRandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
