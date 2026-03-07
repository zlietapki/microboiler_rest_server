package domain

type Repository interface {
	FindAll() ([]User, error)
	Save(*User) error
}
