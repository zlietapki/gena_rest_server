package usecase

import "github.com/zlietapki/gena/internal/domain"

func (uc *Usecase) GetCounter() (domain.Counter, error) {
	return uc.counterRepo.GetNextCounter()
}
