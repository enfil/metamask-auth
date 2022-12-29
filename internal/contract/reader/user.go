package contract

import "github.com/enfil/metamask-auth/internal/domain/user"

type UserReadModel interface {
	GetByAddress(address string) (user.Entity, error)
}
