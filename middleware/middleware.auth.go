package middleware

import (
	"fmt"
	"strings"

	"github.com/ecelliiit/ecell-backend/config"
	"github.com/ecelliiit/ecell-backend/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	secret := strings.Split(auth, " ")[1]
	fmt.Println("secret", secret)

	if secret != config.Cfg.Secret {
		utils.SendResponse(c, 401, "you cannot access the request", nil)
		return
	}

	c.Next()
}
