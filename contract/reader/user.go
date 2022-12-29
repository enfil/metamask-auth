package contract

import (
	"github.com/enfil/metamask-auth/domain/user"
)

type UserReadModel interface {
	GetByAddress(address string) (user.Entity, error)
}
