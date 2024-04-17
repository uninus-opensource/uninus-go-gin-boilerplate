package service

import (
	"errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth/dto"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/hashing"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/jwt"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/timestamp"
)

type authSvc struct {
	userRepo user.UserRepoInterface
	userSvc  user.UserSvcInterface
	hash     hashing.HashInterface
	jwt      jwt.IJwt
}

func NewAuthService(userRepo user.UserRepoInterface, userSvc user.UserSvcInterface, hash hashing.HashInterface, jwt jwt.IJwt) auth.AuthSvcInterface {
	return &authSvc{
		userRepo: userRepo,
		userSvc:  userSvc,
		hash:     hash,
		jwt:      jwt,
	}
}

func (s *authSvc) Register(dto *dto.TypeRegisterRequest) (*entities.MstUser, error) {
	isExistEmail := s.userRepo.FindEmail(dto.Email)
	if isExistEmail != nil {
		log.Error("email is already exist")
		return nil, errors.New("email is already exist")
	}
	hashPass, err := s.hash.GenerateHash(dto.Password)
	if err != nil {
		log.Error("failed hash password")
		return nil, errors.New("failed hash password")
	}

	currentTime, err := timestamp.GetTimestamp()
	if err != nil {
		return nil, err
	}

	newUser := &entities.MstUser{
		Id:        uuid.New().String(),
		Username:  dto.Username,
		Email:     dto.Email,
		Password:  hashPass,
		Phone:     dto.Phone,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	user, err := s.userRepo.Insert(newUser)
	if err != nil {
		log.Error("failed create new user from: service")
		return nil, errors.New("failed create account")
	}
	return user, nil

}

func (s *authSvc) Login(dto *dto.TypeLoginRequest) (*entities.MstUser, string, error) {
	userEmail := s.userRepo.FindEmail(dto.Email)

	isValidPass, err := s.hash.ComparePassword(userEmail.Password, dto.Password)
	if err != nil || !isValidPass {
		log.Error("incorrect password")
		return nil, "", errors.New("incorrect password")
	}

	accessToken, err := s.jwt.GenerateJWT(userEmail.Id, userEmail.Email)
	if err != nil {
		return nil, "", err
	}
	return userEmail, accessToken, nil

}
