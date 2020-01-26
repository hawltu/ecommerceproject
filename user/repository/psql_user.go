package repository

import (
	"database/sql"
	"errors"

	"github.com/hawltu/project1/entity"
)

type UserRepositoryImpl struct {
	conn *sql.DB
}

func NewUserRepositoryImpl(Conn *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}

func (cri *UserRepositoryImpl) Users() ([]entity.User, error) {

	rows, err := cri.conn.Query("SELECT * FROM  user1;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.User{}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.UserName, &user.FName, &user.LName, &user.Password,&user.ID,&user.Email,&user.Mobile,&user.Address,&user.Image,&user.Shopname)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, user)
	}

	return ctgs, nil


}

func (cri *UserRepositoryImpl) User(id int) (entity.User, error) {

	row := cri.conn.QueryRow("SELECT * FROM user1 WHERE id = $1", id)

	c := entity.User{}

	err := row.Scan(&c.UserName, &c.FName, &c.LName, &c.Password,&c.ID,&c.Email,&c.Mobile,&c.Address,&c.Image,&c.Shopname)

	if err != nil {
		return c, err
	}

	return c, nil
}
func (cri *UserRepositoryImpl) UpdateUser(c entity.User) error {

	_, err := cri.conn.Exec("UPDATE user1 SET name=$1,description=$2, image=$3 WHERE id=$4", c.FName, c.UserName, c.Password, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCategory removes a category from a database by its id
func (cri *UserRepositoryImpl) DeleteUser(id int) error {

	_, err := cri.conn.Exec("DELETE FROM user1 WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCategory stores new category information to database

func (cri *UserRepositoryImpl) StoreUser(c entity.User) error {

	_, err := cri.conn.Exec("INSERT INTO user1 (fname,username,password) values($1, $2, $3)", c.FName, c.UserName, c.Password)

	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

