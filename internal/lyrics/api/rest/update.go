package v1

import (
	"net/http"

	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
	"example.com/pkg/request"
	"github.com/gin-gonic/gin"
)

func (c *LyricsController) UpdateLyrics(ctx *gin.Context) {
	request := request.UpdateLyricsRequest{}
	ctx.Bind(&request)
	songId := ctx.Param("songId")

	result := c.service.UpdateLyrics(&dto.UpdateLyricsDto{
		Content: request.Content,
	}, songId)
	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Lyrics not Found"})
		return
	}

	ctx.JSON(http.StatusOK, mapper.MapToResponse(result))
}
