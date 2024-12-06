package repository

import "example.com/internal/lyrics/entity"

type LyricsRepository interface {
	FindAll() []entity.Lyrics
	FindBySongId(songId string) *entity.Lyrics
	Find(id string) *entity.Lyrics
	Save(lyrics *entity.Lyrics)
	Update(lyrics *entity.Lyrics)
}
