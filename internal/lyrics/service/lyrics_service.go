package service

import (
	"time"

	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/entity"
	"example.com/internal/lyrics/mapper"
	"example.com/internal/lyrics/repository"
	"example.com/pkg/constants"
	"github.com/google/uuid"
)

type LyricsService struct {
	repository *repository.LyricsRepository
}

func NewLyricsService() *LyricsService {
	return &LyricsService{
		repository.NewLyricsRepository(),
	}
}

func (s *LyricsService) GetAllLyrics() []*dto.LyricsDto {
	lyricsList := s.repository.FindAll()
	result := []*dto.LyricsDto{}

	for _, lyrics := range lyricsList {
		result = append(result, mapper.MapFromEntity(&lyrics))
	}

	return result
}

func (s *LyricsService) GetLyricsBySongId(songId string) *dto.LyricsDto {
	lyrics := s.repository.FindBySongId(songId)

	if lyrics == nil {
		return nil
	}

	return mapper.MapFromEntity(lyrics)
}

func (s *LyricsService) AddLyrics(dto *dto.CreateLyricsDto) *dto.LyricsDto {
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

func (s *LyricsService) UpdateLyrics(dto *dto.UpdateLyricsDto, songId string) *dto.LyricsDto {
	lyrics := s.repository.FindBySongId(songId)

	if lyrics == nil {
		return nil
	}

	lyrics.Content = dto.Content
	s.repository.Update(lyrics)

	return mapper.MapFromEntity(lyrics)
}
