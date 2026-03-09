package repository

import "github.com/zlietapki/boilerplate/internal/domain"

func (r *Repo) FindAll() ([]domain.User, error) {
	return r.users, nil
}
