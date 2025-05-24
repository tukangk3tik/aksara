package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/utils"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "Bearer"
	AuthorizationPayloadKey = "AuthorizationPayload"
)

func AuthorizeJwt(tokenMaker TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeaderKey)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildErrorResponse("UNAUTHORIZED", utils.ErrorCodeMap["UNAUTHORIZED"], nil))
			return
		}

		strToken := strings.Split(authHeader, " ")
		validateRes, err := tokenMaker.VerifyToken(strToken[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildErrorResponse("UNAUTHORIZED", utils.ErrorCodeMap["UNAUTHORIZED"], nil))
			return
		}

		c.Set("user_id", validateRes.UserId)
		c.Set(AuthorizationPayloadKey, validateRes)
		c.Next()
	}
}
