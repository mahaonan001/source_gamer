package utils

import (
	"math/rand"
	"time"
)

func RandomString(l int, Inner string) string {
	var letters = []byte(Inner)
	var result = make([]byte, l)
	rand.NewSource(time.Now().UnixNano())
	for i := range l {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
