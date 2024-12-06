package lyrics

import (
	"time"

	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/entity"
	"example.com/internal/lyrics/mapper"
	"example.com/pkg/constants"
	"github.com/google/uuid"
)

func (s *lyricsService) AddLyrics(dto *dto.CreateLyricsDto) *dto.LyricsDto {
	lyrics := entity.Lyrics{
		Id:        uuid.NewString(),
		SongId:    dto.SongId,
		Content:   dto.Content,
		CreatedAt: time.Now().Format(constants.DateFormat),
		UpdatedAt: time.Now().Format(constants.DateFormat),
	}

	s.repository.Save(&lyrics)

	return mapper.MapFromEntity(&lyrics)
}
