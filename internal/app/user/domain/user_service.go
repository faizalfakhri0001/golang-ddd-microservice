package domain

type UserService interface {
	Create() error
	GetById(id string) (*User, error)
}
