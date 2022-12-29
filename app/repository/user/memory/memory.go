package memory

import (
	user2 "github.com/enfil/metamask-auth/domain/user"
	"sync"
)

type Repository struct {
	lock  sync.RWMutex
	users map[string]user2.Entity
}

func NewMemoryRepository() *Repository {
	ans := Repository{
		users: make(map[string]user2.Entity),
	}
	return &ans
}

func (repo *Repository) GetByAddress(address string) (user2.Entity, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	u, exists := repo.users[address]
	if !exists {
		return u, user2.ErrUserNotExists
	}
	return u, nil
}

func (repo *Repository) Store(u *user2.Entity) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	if _, exists := repo.users[u.CryptoAddress()]; exists {
		return user2.ErrUserExists
	}
	repo.users[u.CryptoAddress()] = *u
	return nil
}

func (repo *Repository) Update(u *user2.Entity) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	repo.users[u.CryptoAddress()] = *u
	return nil
}
