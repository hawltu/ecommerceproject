package item

import "github.com/hawltu/project1/entity"

type ItemRepository interface {
	Items() ([]entity.Item, error)
	Item(id int) (entity.Item, error)
	UpdateItem(item entity.Item) error
	DeleteItem(id int) error
	StoreItem(item entity.Item) error
}
