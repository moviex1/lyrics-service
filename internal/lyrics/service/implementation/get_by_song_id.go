package lyrics

import (
	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
)

func (s *lyricsService) GetLyricsBySongId(songId string) *dto.LyricsDto {
	lyrics := s.repository.FindBySongId(songId)

	if lyrics == nil {
		return nil
	}

	return mapper.MapFromEntity(lyrics)
}
