// start name:top
package map_repo

// start name:imports type:merge
import "github.com/zlietapki/gena/internal/domain"

// start name:post_imports

var _ domain.ICounterRepo = (*MapRepo)(nil) // compile-time check

type MapRepo struct {
	//start name:repo_struct type:merge
	mem map[string]int
	// start name:post_repo_struct
}

func New() *MapRepo {
	return &MapRepo{
		// start name:repo_init type:merge
		mem: make(map[string]int),
		// start name:post_repo_init
	}
}
