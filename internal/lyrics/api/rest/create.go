package v1

import (
	"net/http"

	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
	"example.com/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (c *LyricsController) CreateLyrics(ctx *gin.Context) {
	request := request.CreateLyricsRequest{}
	ctx.Bind(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	result := c.service.AddLyrics(&dto.CreateLyricsDto{
		SongId:  request.SongId,
		Content: request.Content,
	})

	ctx.JSON(http.StatusOK, mapper.MapToResponse(result))
}
