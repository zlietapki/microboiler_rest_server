package rest_handler

import (
	"github.com/zlietapki/microboiler_rest_server/internal/domain"
	"github.com/zlietapki/microboiler_rest_server/internal/rest_models"
)

func domainUserToRestAPI(u *domain.User) rest_models.User {
	return rest_models.User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func domainUsersToRestAPI(users []domain.User) []rest_models.User {
	result := make([]rest_models.User, len(users))

	for i := range users {
		result[i] = domainUserToRestAPI(&users[i])
	}

	return result
}
