package service

import "example.com/internal/lyrics/dto"

type LyricsService interface {
	UpdateLyrics(dto *dto.UpdateLyricsDto, songId string) *dto.LyricsDto
	AddLyrics(dto *dto.CreateLyricsDto) *dto.LyricsDto
	GetLyricsBySongId(songId string) *dto.LyricsDto
	GetAllLyrics() []*dto.LyricsDto
}
