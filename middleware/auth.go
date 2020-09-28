package middleware

import (
	"easy_curd/model"
	"easy_curd/serializer"
	"easy_curd/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		info, err := util.DecodeJwt(token)
		if err != nil {
			fmt.Println(err)
			c.Next()
			return
		}
		user, err := model.GetUser(info.Subject)
		if err == nil {
			c.Set("user", &user)
			c.Next()
			return
		}
		fmt.Println(err)
		c.Next()
	}
}

// CurrentUser 获取登录用户
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.AuthErr())
		c.Abort()
	}
}
