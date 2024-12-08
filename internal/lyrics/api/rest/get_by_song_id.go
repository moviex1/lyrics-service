package v1

import (
	"net/http"

	"example.com/internal/lyrics/mapper"
	"github.com/gin-gonic/gin"
)

func (c *LyricsController) GetLyricsBySongId(ctx *gin.Context) {
	songId := ctx.Param("songId")

	lyricsDto := c.service.GetLyricsBySongId(songId)
	if lyricsDto == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Lyrics not Found"})
		return
	}

	ctx.JSON(http.StatusOK, mapper.MapToResponse(lyricsDto))
}
