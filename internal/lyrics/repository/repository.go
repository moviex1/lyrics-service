package repository

import (
	"time"

	postgres "example.com/db"
	"example.com/internal/lyrics/entity"
	"example.com/pkg/constants"
	"github.com/jmoiron/sqlx"
)

type LyricsRepository struct {
	conn *sqlx.DB
}

func NewLyricsRepository() *LyricsRepository {
	return &LyricsRepository{
		conn: postgres.GetConnection(),
	}
}

func (repo *LyricsRepository) FindAll() []entity.Lyrics {
	result := []entity.Lyrics{}
	repo.conn.Select(&result, "SELECT id, song_id, content, created_at, updated_at FROM lyrics")

	return result
}

func (repo *LyricsRepository) FindBySongId(songId string) *entity.Lyrics {
	var lyrics entity.Lyrics

	if err := repo.conn.Get(&lyrics, "SELECT id, song_id, content, created_at, updated_at FROM lyrics WHERE song_id = $1", songId); err != nil {
		return nil
	}

	return &lyrics
}

func (repo *LyricsRepository) Save(lyrics *entity.Lyrics) {
	repo.conn.MustExec("INSERT INTO lyrics (id, song_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", lyrics.Id, lyrics.SongId, lyrics.Content, lyrics.CreatedAt, lyrics.UpdatedAt)
}

func (repo *LyricsRepository) Update(lyrics *entity.Lyrics) {
	repo.conn.MustExec("UPDATE lyrics SET content = $1, updated_at = $2 WHERE id = $3", lyrics.Content, time.Now().Format(constants.DateFormat), lyrics.Id)
}

func (repo *LyricsRepository) Find(id string) *entity.Lyrics {
	lyrics := entity.Lyrics{}
	if err := repo.conn.Get(&lyrics, "SELECT id, song_id, content, created_at, updated_at FROM lyrics WHERE id = $1", id); err != nil {
		return nil
	}

	return &lyrics
}
