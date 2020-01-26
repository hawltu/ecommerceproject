package repository

import (
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/item"
	"github.com/jinzhu/gorm"
)

// PostGormRepo implements the post.PostRepository interface
type ItemGormRepo struct {
	conn *gorm.DB
}

// NewPostGormRepo will create a new object of PostGormRepo
func NewItemGormRepo(db *gorm.DB) item.ItemRepository{
	return &ItemGormRepo{conn: db}
}

// Posts returns all posts stored in the database
func (pRepo *ItemGormRepo) Items() ([]entity.Item, error) {
	posts := []entity.Item{}
	errs := pRepo.conn.Find(&posts).GetErrors()
	if len(errs) > 0 {
		return nil, errs[1]
	}

	return posts, nil
}

// Post retrieve a post from the database by its id
func (pRepo *ItemGormRepo) Item(id int) (entity.Item, error) {
	post := entity.Item{}
	errs := pRepo.conn.First(&post, id).GetErrors()
	if len(errs) > 0 {
		return post, errs[1]
	}
	return post, nil
}

// UpdatePost updates a given post in the database
func (pRepo *ItemGormRepo) UpdateItem(post entity.Item) (error) {
	pst := post
	errs := pRepo.conn.Save(pst).GetErrors()
	if len(errs) > 0 {
		return errs[1]
	}
	return nil
}

// DeletePost deletes a given post from the database
func (pRepo *ItemGormRepo) DeleteItem(id int) (error) {
	post, errs := pRepo.Item(id)
	if errs != nil{
		return nil
	}
	idd := post.ID
	err := pRepo.conn.Delete(post, idd).GetErrors()
	if err != nil {
		return nil
	}
	return nil
}

// StorePost stores a given post in the database
func (pRepo *ItemGormRepo) StoreItem(post entity.Item) (error) {
	pst := post
	errs := pRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0 {
		return errs[1]
	}
	return nil
}
