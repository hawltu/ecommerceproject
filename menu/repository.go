package menu

import "github.com/hawltu/project1/entity"

// CategoryRepository specifies menu category related database operations
type CategoryRepository interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
