package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user"
)

type userSvc struct {
	repo user.UserRepoInterface
}

func NewUserService(repo user.UserRepoInterface) user.UserSvcInterface {
	return &userSvc{
		repo: repo,
	}
}

func (s *userSvc) GetId(Id string)(*entities.MstUser, error) {
	user, err := s.repo.FindId(Id)
	if err != nil {
		log.Error("failed get id from service")
		return nil, errors.New("failed get id from service")
	}
	return user, nil
}