package access

import (
	"math/rand"
	"time"
)

type AccessToken string

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-="
const length int = 32

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateAccessToken() AccessToken {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return AccessToken(b)
}
