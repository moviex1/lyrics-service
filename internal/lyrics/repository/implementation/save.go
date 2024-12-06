package lyrics

import "example.com/internal/lyrics/entity"

func (repo *lyricsRepository) Save(lyrics *entity.Lyrics) {
	repo.conn.MustExec("INSERT INTO lyrics (id, song_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", lyrics.Id, lyrics.SongId, lyrics.Content, lyrics.CreatedAt, lyrics.UpdatedAt)
}
