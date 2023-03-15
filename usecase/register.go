package usecase

import (
	"github.com/enfil/metamask-auth/contract/service"
	"github.com/enfil/metamask-auth/domain/user"
	"github.com/enfil/metamask-auth/usecase/command"
	"github.com/google/uuid"
	"strings"
)

type Registrar struct {
	Repo          user.Repository
	NonceProvider contract.NonceProvider
}

func (r *Registrar) Handle(c command.Register) error {
	nonce, err := r.NonceProvider.GenerateNonce()
	if err != nil {
		return err
	}
	u, err := user.New(uuid.New(), strings.ToLower(c.CryptoAddress), nonce, "")
	if err != nil {
		return err
	}
	err = r.Repo.Store(&u)
	if err != nil {
		return err
	}
	return nil
}
