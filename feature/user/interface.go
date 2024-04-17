package user

import "github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"

type UserRepoInterface interface {
	FindId(id string) (*entities.MstUser, error)
	Insert(data *entities.MstUser) (*entities.MstUser, error)
	FindEmail(email string) (*entities.MstUser)
}

type UserSvcInterface interface {
	GetId(id string) (*entities.MstUser, error)
}
