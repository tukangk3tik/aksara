package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/utils"
)

func AuthorizeJwt(tokenMaker TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.BuildErrorResponse("UNAUTHORIZED", utils.ErrorCodeMap["UNAUTHORIZED"], nil))
			return
		}

		strToken := strings.Split(authHeader, " ")
		validateRes, err := tokenMaker.VerifyToken(strToken[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildErrorResponse("UNAUTHORIZED", utils.ErrorCodeMap["UNAUTHORIZED"], nil))
			return
		}

		c.Set("user_id", validateRes.UserId)
		c.Next()
	}
}
