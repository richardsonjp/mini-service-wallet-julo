package middlewares

import (
	"github.com/gin-gonic/gin"
	constUtil "go-skeleton/pkg/utils/constant"
	"go-skeleton/pkg/utils/errors"
	"strings"
)

// Generic header static API key validator middleware
func (m *MiddlewareAccess) AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
				{Key: "Authorization", Value: "Authorization Header not found"},
			}))
			return
		}

		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			if token[0] != "Bearer" {
				errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
					{Key: "Authorization", Value: "Bearer not found"},
				}))
				return
			} else {
				errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
					{Key: "Authorization", Value: "Token not found"},
				}))
				return
			}
		}

		SID := token[1]
		customerID, err := m.credentialService.AuthenticateToken(c, SID)
		if err != nil {
			errors.ErrorCode(c, errors.CLIENT_AUTH_ERROR, errors.SetDetails([]constUtil.ErrorDetails{
				{Key: "Authorization", Value: "Token not found"},
			}))
		}

		c.Set("customer_id", customerID)

		c.Next()
	}
}
