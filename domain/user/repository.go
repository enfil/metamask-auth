package user

type Repository interface {
	GetByAddress(address string) (Entity, error)
	Store(u *Entity) error
	Update(u *Entity) error
}
