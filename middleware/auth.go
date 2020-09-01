package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"easy_curd/model"
	"easy_curd/serializer"
	"easy_curd/util"
)

// CurrentUser 获取登录用户
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		info, _ := util.DecodeJwt(token)
		if info != nil {
			user, err := model.GetUser(info.Subject)
			if err == nil {
				c.Set("user", &user)
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()

	}
}
