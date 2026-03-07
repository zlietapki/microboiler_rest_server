package usecase

import "github.com/zlietapki/microboiler_rest_server/internal/domain"

type UserUseCase struct {
	repo domain.Repository
}

func New(repo domain.Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetAll() ([]domain.User, error) {
	return uc.repo.FindAll()
}

func (uc *UserUseCase) Create(name, email string) (*domain.User, error) {
	u := &domain.User{Name: name, Email: email}
	if err := uc.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}
