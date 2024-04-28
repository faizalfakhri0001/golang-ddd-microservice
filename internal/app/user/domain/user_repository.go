package domain


type UserRepository interface {
	Create() error
	GetById(id string) (*User, error)
}
