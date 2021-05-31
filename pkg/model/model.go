package model

// User holds data associated with registered user
type User struct {
	IDToken string `form:"idtoken" json:"idtoken" binding:"required"`
}
