// start name:top
package domain

// start name:usecase
type IUsecase interface {
	//start name:methods type:merge
	GetCounter() (Counter, error)
	//start name:post_methods
}
