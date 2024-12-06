package lyrics

import (
	v1 "example.com/api/v1"
	r "example.com/internal/lyrics/repository/implementation"
	s "example.com/internal/lyrics/service/implementation"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	lyricsController := v1.NewLyricsController(
		s.NewLyricsService(
			r.NewLyricsRepository(),
		),
	)

	c.GET("/api/v1/lyrics", lyricsController.GetAllLyrics)
	c.GET("/api/v1/songs/:songId/lyrics", lyricsController.GetLyricsBySongId)
	c.POST("/api/v1/lyrics", lyricsController.CreateLyrics)
	c.PUT("/api/v1/songs/:songId/lyrics", lyricsController.UpdateLyrics)
}
