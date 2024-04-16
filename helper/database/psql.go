package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/config"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionDB(cnf config.APPConfig) *gorm.DB {
	psqlConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cnf.DBConf.DbHost, cnf.DBConf.DbUser, cnf.DBConf.DbPass, cnf.DBConf.DbName, cnf.DBConf.DbPort)
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{})
	if err != nil {
		log.Error("error open connection")
		return nil
	}
	log.Info("database connected")
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entities.MstUser{}); err != nil {
		log.Fatal("error migrate table")
	}
	log.Println("migrate succesfully")
}
