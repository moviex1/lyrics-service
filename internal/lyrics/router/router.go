package router

import (
	"example.com/internal/lyrics/router/lyrics"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	lyrics.RegisterRoutes(c)
}
