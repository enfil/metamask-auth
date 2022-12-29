package contract

import (
	"github.com/enfil/metamask-auth/pkg/domain/user"
)

type UserReadModel interface {
	GetByAddress(address string) (user.Entity, error)
}
