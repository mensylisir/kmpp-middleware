package aop

import (
	"github.com/gin-gonic/gin"
	"github.com/mensylisir/kmpp-middleware/src/util/jwt"
	"github.com/toolkits/pkg/ginx"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		profile, err := jwt.ExtractTokenMetadata(c.Request)
		if err != nil {
			ginx.Bomb(http.StatusUnauthorized, "unauthorized")
		}

		c.Set("user", profile)
		c.Next()

	}
}
