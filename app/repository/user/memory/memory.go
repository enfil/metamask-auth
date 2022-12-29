package memory

import (
	"github.com/enfil/metamask-auth/pkg/domain/user"
	"sync"
)

type Repository struct {
	lock  sync.RWMutex
	users map[string]user.Entity
}

func NewMemoryRepository() *Repository {
	ans := Repository{
		users: make(map[string]user.Entity),
	}
	return &ans
}

func (repo *Repository) GetByAddress(address string) (user.Entity, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	u, exists := repo.users[address]
	if !exists {
		return u, user.ErrUserNotExists
	}
	return u, nil
}

func (repo *Repository) Store(u *user.Entity) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	if _, exists := repo.users[u.CryptoAddress()]; exists {
		return user.ErrUserExists
	}
	repo.users[u.CryptoAddress()] = *u
	return nil
}

func (repo *Repository) Update(u *user.Entity) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	repo.users[u.CryptoAddress()] = *u
	return nil
}
