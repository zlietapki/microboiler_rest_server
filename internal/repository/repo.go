package repository

import "github.com/zlietapki/microboiler_rest_server/internal/domain"

type Repo struct {
	users  []domain.User
	nextID int
}

func New() *Repo {
	return &Repo{nextID: 1}
}

func (r *Repo) FindAll() ([]domain.User, error) {
	return r.users, nil
}

func (r *Repo) Save(u *domain.User) error {
	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *u)
	return nil
}
