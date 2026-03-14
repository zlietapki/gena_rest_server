package domain

type ICounterRepo interface {
	GetNextCounter() (Counter, error)
}
