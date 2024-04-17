package main

import (
	"fmt"

	uRepo "github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user/repository"
	uSvc "github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user/service"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/hashing"
	jwt2 "github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/jwt"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/routes"

	"github.com/gin-gonic/gin"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/config"
	aHandler "github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth/handler"
	aSvc "github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/auth/service"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/database"
)

func main() {
	app := gin.Default()

	var loadConfig = config.RunConfig()

	db := database.NewConnectionDB(*loadConfig)
	database.Migrate(db)
	jwt := jwt2.NewJWT(loadConfig.Secret)
	hash := hashing.NewHash()

	userRepo := uRepo.NewUserRepository(db)
	userSvc := uSvc.NewUserService(userRepo)

	authSvc := aSvc.NewAuthService(userRepo, userSvc, hash, jwt)
	authHand := aHandler.NewAuthHandler(authSvc)

	app.Use(corsMiddleware())
	
	routes.AuthRoute(app, authHand)
	addr := fmt.Sprintf(":%d", loadConfig.AppPort)
	app.Run(addr)

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
