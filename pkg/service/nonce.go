package service

import (
	"crypto/rand"
	"math/big"
	"sync"
)

type Nonce struct {
	max  *big.Int
	once sync.Once
}

func (n *Nonce) GenerateNonce() (string, error) {
	n.once.Do(func() {
		n.max = new(big.Int)
		n.max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(n.max, big.NewInt(1))
	})
	nonce, err := rand.Int(rand.Reader, n.max)
	if err != nil {
		return "", err
	}
	return nonce.Text(10), nil
}
