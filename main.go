package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/config"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/database"
)

func main() {
	app := gin.Default()

	var loadConfig = config.RunConfig()

	db := database.NewConnectionDB(*loadConfig)
	database.Migrate(db)

	app.Use(corsMiddleware())

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
