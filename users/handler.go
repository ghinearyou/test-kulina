package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func (h Handler) handleCreateUser(c *gin.Context) {
	var params User
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := h.service.CreateUser(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, Response("Success"))
	return
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func Response(data interface{}) gin.H {
	return gin.H{
		"data":  data,
		"error": nil,
	}
}
