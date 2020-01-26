package item

import "github.com/hawltu/project1/entity"


type ItemService interface {
	Items() ([]entity.Item, error)
	Item(id int) (entity.Item, error)
	UpdateItem(user entity.Item) error
	DeleteItem(id int) error
	StoreItem(item entity.Item) error
}

