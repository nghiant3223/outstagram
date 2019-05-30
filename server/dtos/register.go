package dtos

type RegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Fullname string `form:"fullname" json:"fullname" binding:"required"`
}
