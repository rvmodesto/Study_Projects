package middleware

import (
	token "e-commerce/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func (c *gin.Context) {
		ClientToken := c.Request.Header.Get("token") //getting token -> postman
		if ClientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"No authorization header provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)//token - reference: token "e-commerce/tokens"
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next() //whenever a api is called it sees if the token is valid or not, so it can proceed to the 'next' thing
	}
}