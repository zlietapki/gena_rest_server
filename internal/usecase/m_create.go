package usecase

import "github.com/zlietapki/boilerplate/internal/domain"

func (uc *Usecase) Create(name, email string) (*domain.User, error) {
	u := &domain.User{Name: name, Email: email}
	if err := uc.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}
