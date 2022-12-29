package reader

import (
	contract "github.com/enfil/metamask-auth/internal/contract/reader"
	"github.com/enfil/metamask-auth/internal/domain/user"
	"strings"
)

type User struct {
	ReadModel contract.UserReadModel
}

func (u *User) ByAddress(address string) (user.Entity, error) {
	return u.ReadModel.GetByAddress(strings.ToLower(address))
}
