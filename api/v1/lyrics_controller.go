package v1

import (
	"net/http"

	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
	"example.com/internal/lyrics/service"
	"example.com/pkg/request"
	"example.com/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LyricsController struct {
	service service.LyricsService
}

func NewLyricsController(
	service service.LyricsService,
) *LyricsController {
	return &LyricsController{
		service,
	}
}

func (c *LyricsController) GetAllLyrics(ctx *gin.Context) {
	lyricsDtos := c.service.GetAllLyrics()

	result := []*response.LyricsResponse{}
	for _, lyricsDto := range lyricsDtos {
		result = append(result, mapper.MapToResponse(lyricsDto))
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *LyricsController) GetLyricsBySongId(ctx *gin.Context) {
	songId := ctx.Param("songId")

	lyricsDto := c.service.GetLyricsBySongId(songId)
	if lyricsDto == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Lyrics not Found"})
		return
	}

	ctx.JSON(http.StatusOK, mapper.MapToResponse(lyricsDto))
}

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

func (c *LyricsController) UpdateLyrics(ctx *gin.Context) {
	request := request.UpdateLyricsRequest{}
	ctx.Bind(&request)
	songId := ctx.Param("songId")

	result := c.service.UpdateLyrics(&dto.UpdateLyricsDto{
		Content: request.Content,
	}, songId)

	ctx.JSON(http.StatusOK, mapper.MapToResponse(result))
}
