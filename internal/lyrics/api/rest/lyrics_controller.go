package v1

import (
	"example.com/internal/lyrics/service"
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
