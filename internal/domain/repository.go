// start name:all
package domain

type IRepository interface {
	//start name:repo_methods type:add
	FindAll() ([]User, error)
	Save(*User) error
	//	start name:post_repo_methods
}
