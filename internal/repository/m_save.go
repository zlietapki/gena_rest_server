package repository

import "github.com/zlietapki/boilerplate/internal/domain"

func (r *Repo) Save(u *domain.User) error {
	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *u)
	return nil
}
