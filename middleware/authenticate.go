package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		fmt.Println(cookie)
		if err != nil {
			cookie = "NotSet"
		}

		c.Next()
	}
}
