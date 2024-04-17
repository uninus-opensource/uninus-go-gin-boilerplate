package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepoInterface {
	return &userRepo{db: db}
}

func (r *userRepo) FindId(Id string) (*entities.MstUser, error) {
	var user *entities.MstUser
	if err := r.db.Where("id = ? AND deleted_at IS NULL", Id).First(&user).Error; err != nil {
		log.Error("failed find user by id")
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Insert(user *entities.MstUser) (*entities.MstUser, error) {
	if err := r.db.Create(user).Error; err != nil {
		log.Error("failed create new data user from: repository")
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindEmail(email string) *entities.MstUser {
	var user *entities.MstUser
	if err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		log.Error("failed get email user from: repository")
		return nil
	}
	return user
}


