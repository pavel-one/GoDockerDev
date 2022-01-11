package requests

type Register struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}
