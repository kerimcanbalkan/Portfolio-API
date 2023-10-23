package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectController(c *gin.Context) {
	c.String(http.StatusOK, "Hello From Projects")
}
