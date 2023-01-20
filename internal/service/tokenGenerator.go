package service

import (
	"crypto/rand"
	"math/big"
)

const (
	literals   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	linkLength = 10
)

// TokenGenerator generates new token for shortened link
func TokenGenerator() (string, error) {
	ret := make([]byte, linkLength)
	for i := 0; i < linkLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(literals))))
		if err != nil {
			return "", err
		}

		ret[i] = literals[num.Int64()]
	}

	return string(ret), nil
}
