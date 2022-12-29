package reader

import (
	"github.com/enfil/metamask-auth/pkg/contract/reader"
	"github.com/enfil/metamask-auth/pkg/domain/user"
	"strings"
)

type User struct {
	ReadModel contract.UserReadModel
}

func (u *User) ByAddress(address string) (user.Entity, error) {
	return u.ReadModel.GetByAddress(strings.ToLower(address))
}
