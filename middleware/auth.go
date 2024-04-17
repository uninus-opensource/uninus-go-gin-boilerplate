package middleware

import (
	"strconv"
	"strings"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	jwtMapClaims "github.com/golang-jwt/jwt"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/feature/user"
	"github.com/uninus-opensource/uninus-go-gin-boilerplate/helper/jwt"
)

func Protected(jwtService jwt.IJwt, userSvc user.UserSvcInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: invalid token"})
			c.Abort()
			return
		}

		claim, ok := token.Claims.(jwtMapClaims.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: invalid token"})
			c.Abort()
			return
		}

		userIDStr, ok := claim["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: invalid user ID"})
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		user, err := userSvc.GetId(strconv.Itoa(userID))
		if err != nil {
			log.Printf("Error retrieving user: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied: user not found"})
			c.Abort()
			return
		}

		c.Set("CurrentUser", user)

	}
}
