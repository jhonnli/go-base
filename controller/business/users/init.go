package users

import "github.com/gin-gonic/gin"

func Init(engin *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitUsers(engin))
	return data
}
