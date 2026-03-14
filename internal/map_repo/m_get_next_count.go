package map_repo

import "github.com/zlietapki/gena/internal/domain"

func (r *MapRepo) GetNextCounter() (domain.Counter, error) {
	r.mem["count"]++

	return domain.Counter{
		Value: r.mem["count"],
	}, nil
}
