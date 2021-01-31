package model

	import "gopkg.in/go-playground/validator.v9"
// User holds data associated with registered user
type User struct {
	Email    string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
func (u *User)IsValid() (bool,error) {
	v:=validator.New()
	if err:=v.Struct(u);err!=nil {
		return false ,err
	}
	return true,nil
}
