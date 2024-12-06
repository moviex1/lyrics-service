package lyrics

import (
	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
)

func (s *lyricsService) GetAllLyrics() []*dto.LyricsDto {
	lyricsList := s.repository.FindAll()
	result := []*dto.LyricsDto{}

	for _, lyrics := range lyricsList {
		result = append(result, mapper.MapFromEntity(&lyrics))
	}

	return result
}
