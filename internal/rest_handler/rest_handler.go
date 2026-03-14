package rest_handler

import "github.com/zlietapki/gena/internal/usecase"

type UserHandler struct {
	uc *usecase.Usecase
}

func New(uc *usecase.Usecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}
