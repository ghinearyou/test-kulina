package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// Register register routes for health
func Register(router *gin.Engine, service *Service) {
	handler := &Handler{
		service: service,
	}
	v1 := router.Group("/v1")
	{
		v1.POST("/create", handler.handleCreateUser)
		v1.GET("/split/:number",func (c *gin.Context) {
		value := c.Params.ByName("number")
		var temp string
		for i:= 0; i < len(value); i++ {
			temp = string(value[i])
			fmt.Println(temp + strings.Repeat("0", len(value) - (i + 1)))
		}
		//fmt.Println(str2)

	})
	}
}
