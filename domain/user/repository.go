package user

type Repository interface {
	// Store /*Fetch(cursor string, num int64) (res []Entity, nextCursor string, err error)

	GetByAddress(address string) (Entity, error)
	Store(u *Entity) error
	Update(u *Entity) error
	// Get(uuid uuid.UUID)
}
