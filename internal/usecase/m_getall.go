package usecase

import "github.com/zlietapki/boilerplate/internal/domain"

func (uc *Usecase) GetAll() ([]domain.User, error) {
	return uc.repo.FindAll()
}
