package entity

type UserSession struct {
	UUID   string
	UserID string
}
type Item struct {
	id          int
	name        string
	catagory    string
	subcatagory string
	price       float32
	quantity    int
	image       string
}
