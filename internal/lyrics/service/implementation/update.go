package lyrics

import (
	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/mapper"
)

func (s *lyricsService) UpdateLyrics(dto *dto.UpdateLyricsDto, songId string) *dto.LyricsDto {
	lyrics := s.repository.FindBySongId(songId)

	if lyrics == nil {
		return nil
	}

	lyrics.Content = dto.Content
	s.repository.Update(lyrics)

	return mapper.MapFromEntity(lyrics)
}
