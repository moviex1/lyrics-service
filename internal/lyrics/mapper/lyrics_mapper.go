package mapper

import (
	"example.com/internal/lyrics/dto"
	"example.com/internal/lyrics/entity"
	"example.com/pkg/response"
)

func MapFromEntity(lyrics *entity.Lyrics) *dto.LyricsDto {
	return &dto.LyricsDto{
		Id:        lyrics.Id,
		SongId:    lyrics.SongId,
		Content:   lyrics.Content,
		CreatedAt: lyrics.CreatedAt,
		UpdatedAt: lyrics.UpdatedAt,
	}
}

func MapToResponse(lyrics *dto.LyricsDto) *response.LyricsResponse {
	return &response.LyricsResponse{
		Id:        lyrics.Id,
		SongId:    lyrics.SongId,
		Content:   lyrics.Content,
		CreatedAt: lyrics.CreatedAt,
		UpdatedAt: lyrics.UpdatedAt,
	}
}
