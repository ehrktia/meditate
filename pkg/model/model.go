package model

// User holds data associated with registered user
type User struct {
	Email    string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
