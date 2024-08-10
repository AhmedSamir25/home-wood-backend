package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phonenumber"`
	Email    string `json:"email"`
}
