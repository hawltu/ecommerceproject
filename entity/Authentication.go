package entity

type Authenticate struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
