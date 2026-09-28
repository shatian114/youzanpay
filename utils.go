package youzan

import (
	"math/rand/v2"
)

const randSalt = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const randSaltLen = uint(len(randSalt))

func RandStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randSalt[rand.UintN(randSaltLen)]
	}
	return string(b)
}
