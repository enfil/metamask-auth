package reader

import (
	"github.com/enfil/metamask-auth/contract/reader"
	"github.com/enfil/metamask-auth/domain/user"
	"strings"
)

type User struct {
	ReadModel contract.UserReadModel
}

func (u *User) ByAddress(address string) (user.Entity, error) {
	return u.ReadModel.GetByAddress(strings.ToLower(address))
}
