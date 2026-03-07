package rest_handler

import "github.com/zlietapki/microboiler_rest_server/internal/usecase"

type UserHandler struct {
	uc *usecase.UserUseCase
}

func New(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}
