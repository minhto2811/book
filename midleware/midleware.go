package midleware

import (
	"book/models"
	security "book/security"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort()
			return
		}
		
		token = strings.Replace(token, "Bearer ", "", 1)

		claims, err := security.ParseAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort()
		}

		iss, err := claims.GetIssuer()
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort() 
			return
		}

		if iss != os.Getenv("APP_NAME") {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort() 
			return
		}

		exp, err := claims.GetExpirationTime()
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort() 
			return
		}

		if exp == nil || exp.UTC().Unix() < time.Now().UTC().Unix() {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token expired"))
			c.Abort() 
			return
		}

		userId, err := claims.GetSubject()
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse("Unauthorized"))
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next() 
	}
}
