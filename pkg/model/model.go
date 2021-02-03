package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// User holds data associated with registered user
type User struct {
	Email    string `form:"username" json:"username" binding:"required" validate:"required,email"`
	Password string `form:"password" json:"password" binding:"required" validate:"required,alphanum"`
}

func (u *User) IsValid() error {
	v := validator.New()
	if err := v.Struct(u); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("error validate fields: %v with value: %v", e.Field(), e.Value())
		}
	}
	return nil
}
