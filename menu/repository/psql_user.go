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
<<<<<<< HEAD
		err = rows.Scan(&user.username, &user.Fname, &user.Lname, &user.Password, &user.shopname, &user.address, &user.mobile)
=======
		err = rows.Scan(&user.Username, &user.Fname, &user.Lname, &user.Password)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
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

<<<<<<< HEAD
	err := row.Scan(&c.username, &c.Fname,&c.Lname,&c.password, &c.email,&c.shopname,&c.address,&c.mobile)
=======
	err := row.Scan(&c.Fname, &c.Lname, &c.Username, &c.Password)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
	if err != nil {
		return c, err
	}

	return c, nil
}

<<<<<<< HEAD
func (cri *ItemRepositoryImpl) UpdateUser(c entity.User) error {

	_, err := cri.conn.Exec("UPDATE user1 SET shopname=$1, address=$2 WHERE id=$3", c.shopname, c.address, c.username)
=======
func (cri *UserRepositoryImpl) UpdateUser(c entity.UserUser) error {

	_, err := cri.conn.Exec("UPDATE user1 SET name=$1,description=$2, image=$3 WHERE id=$4", c.fname, c.username, c.password, c.ID)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
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
<<<<<<< HEAD
func (cri *UserRepositoryImpl) StoreUser(c *entity.User) error {

	_, err := cri.conn.Exec("INSERT INTO user1 (fname,username,password,email,shopname,adress) values($1, $2, $3,$4,$5,$6)", &c.fullname, &c.username, &c.password, &c.email,&c.shopname,&c.address)
=======
func (cri *UserRepositoryImpl) StoreUser(c entity.User) error {

	_, err := cri.conn.Exec("INSERT INTO user1 (fname,username,password) values($1, $2, $3)", c.Fname, c.Username, c.Password)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
<<<<<<< HEAD

=======
	
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
}
