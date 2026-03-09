// start name:top
package domain

// start name:imports type:merge
// start name:usecase
type IUsecase interface {
	//start name:methods type:merge
	GetAll() ([]User, error)
	Create(name, email string) (*User, error)
	//start name:post_methods
}
