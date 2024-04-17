package dto

import "github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"

type TypeLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"access_token"`
}

func LoginResponse(user *entities.MstUser, token string) *TypeLoginResponse {
	userFormatter := &TypeLoginResponse{}
	userFormatter.Username = user.Username
	userFormatter.Email = user.Email
	userFormatter.Phone = user.Phone
	userFormatter.Token = token

	return userFormatter
}
