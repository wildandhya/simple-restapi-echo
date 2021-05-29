package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required,gte=0,lte=10"`
}
