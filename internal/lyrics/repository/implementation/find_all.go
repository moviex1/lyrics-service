package lyrics

import "example.com/internal/lyrics/entity"

func (repo *lyricsRepository) FindAll() []entity.Lyrics {
	result := []entity.Lyrics{}
	repo.conn.Select(&result, "SELECT id, song_id, content, created_at, updated_at FROM lyrics")

	return result
}
