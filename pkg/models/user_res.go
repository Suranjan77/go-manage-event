package models

import "github.com/Suranjan77/go-manage-event/pkg/domain"

type UserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func UserResponseFromUser(u *domain.User) *UserResponse {
	usrRes := UserResponse{}
	usrRes.Email = u.Email
	usrRes.FirstName = u.FirstName
	usrRes.LastName = u.LastName
	return &usrRes
}
