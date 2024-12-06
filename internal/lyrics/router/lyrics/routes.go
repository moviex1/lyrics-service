package lyrics

import (
	v1 "example.com/api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	lyricsController := v1.NewLyricsController()

	c.GET("/api/v1/lyrics", lyricsController.GetAllLyrics)
	c.GET("/api/v1/songs/:songId/lyrics", lyricsController.GetLyricsBySongId)
	c.POST("/api/v1/lyrics", lyricsController.CreateLyrics)
	c.PUT("/api/v1/songs/:songId/lyrics", lyricsController.UpdateLyrics)
}
