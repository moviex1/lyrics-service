package lyrics

import (
	"example.com/internal/lyrics/repository"
	"example.com/internal/lyrics/service"
)

var _ service.LyricsService = (*lyricsService)(nil)

type lyricsService struct {
	repository *repository.LyricsRepository
}

func NewLyricsService() *lyricsService {
	return &lyricsService{
		repository.NewLyricsRepository(),
	}
}
