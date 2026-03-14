package rest_handler

import (
	"github.com/zlietapki/gena/internal/domain"
	"github.com/zlietapki/gena/internal/rest_models"
)

func domainCounterToRestAPI(u domain.Counter) rest_models.Counter {
	return rest_models.Counter{
		Value: u.Value,
	}
}
