package usecase

import (
	"github.com/enfil/metamask-auth/contract/service"
	user2 "github.com/enfil/metamask-auth/domain/user"
	"github.com/enfil/metamask-auth/usecase/command"
	"strings"
)

type Registrar struct {
	Repo          user2.Repository
	NonceProvider contract.NonceProvider
}

func (r *Registrar) Handle(c command.Register) error {
	nonce, err := r.NonceProvider.GenerateNonce()
	if err != nil {
		return err
	}
	u, err := user2.New(strings.ToLower(c.CryptoAddress), nonce, "")
	if err != nil {
		return err
	}
	err = r.Repo.Store(&u)
	if err != nil {
		return err
	}
	return nil
}
