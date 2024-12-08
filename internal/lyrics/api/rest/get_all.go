package v1

import (
	"net/http"

	"example.com/internal/lyrics/mapper"
	"example.com/pkg/response"
	"github.com/gin-gonic/gin"
)

func (c *LyricsController) GetAllLyrics(ctx *gin.Context) {
	lyricsDtos := c.service.GetAllLyrics()

	result := []*response.LyricsResponse{}
	for _, lyricsDto := range lyricsDtos {
		result = append(result, mapper.MapToResponse(lyricsDto))
	}

	ctx.JSON(http.StatusOK, result)
}
