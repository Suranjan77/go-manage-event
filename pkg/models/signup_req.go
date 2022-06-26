package models

import (
	"bytes"

	"github.com/Suranjan77/go-manage-event/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

func (s SignUpRequest) ToUser() *domain.User {
	var user domain.User
	user.Email = s.Email
	user.FirstName = s.FirstName
	user.LastName = s.LastName
	enc, _ := bcrypt.GenerateFromPassword([]byte(s.Password), 16)
	user.Password = bytes.NewBuffer(enc).String()
	return &user
}
