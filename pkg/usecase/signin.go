package usecase

import (
	contract2 "github.com/enfil/metamask-auth/pkg/contract/service"
	"github.com/enfil/metamask-auth/pkg/domain/user"
	"github.com/enfil/metamask-auth/pkg/usecase/command"
	"strings"
)

type SignIn struct {
	Repo          user.Repository
	NonceProvider contract2.NonceProvider
	SignProvider  contract2.SignProvider
}

func (r SignIn) Handle(c command.SignIn) error {
	u, err := r.Repo.GetByAddress(c.CryptoAddress)
	if err != nil {
		return err
	}
	if u.Nonce() != c.Nonce {
		return user.ErrAuthError
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
