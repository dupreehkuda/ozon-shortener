package service

import (
	"crypto/rand"
	"math"
	"math/big"
)

const (
	literals   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	linkLength = 10
)

// TokenGenerator generates new token for shortened link
func TokenGenerator() (string, error) {
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow(float64(len(literals)), 10))))
	if err != nil {
		return "", err
	}

	num := rnd.Int64()

	res := make([]byte, linkLength)

	for i := 0; i < linkLength; i++ {
		if num > 0 {
			res[i] = literals[num%int64(len(literals))]
			num /= int64(len(literals))
		} else {
			res[i] = literals[0]
		}
	}

	return string(res), nil
}
