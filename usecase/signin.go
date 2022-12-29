package usecase

import (
	"github.com/enfil/metamask-auth/contract/service"
	user2 "github.com/enfil/metamask-auth/domain/user"
	"github.com/enfil/metamask-auth/usecase/command"
	"strings"
)

type SignIn struct {
	Repo          user2.Repository
	NonceProvider contract.NonceProvider
	SignProvider  contract.SignProvider
}

func (r SignIn) Handle(c command.SignIn) error {
	u, err := r.Repo.GetByAddress(c.CryptoAddress)
	if err != nil {
		return err
	}
	if u.Nonce() != c.Nonce {
		return user2.ErrAuthError
	}

	err = r.SignProvider.Check(u.CryptoAddress(), c.Nonce, c.Sig)
	if err != nil {
		return err
	}

	// update the nonce here so that the signature cannot be resused
	nonce, err := r.NonceProvider.GenerateNonce()
	if err != nil {
		return err
	}

	err = u.Edit(strings.ToLower(c.CryptoAddress), nonce, c.Sig)
	if err != nil {
		return err
	}

	err = r.Repo.Update(&u)
	if err != nil {
		return err
	}

	return nil
}
