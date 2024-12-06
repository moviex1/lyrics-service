package lyrics

import "example.com/internal/lyrics/entity"

func (repo *lyricsRepository) Find(id string) *entity.Lyrics {
	lyrics := entity.Lyrics{}
	if err := repo.conn.Get(&lyrics, "SELECT id, song_id, content, created_at, updated_at FROM lyrics WHERE id = $1", id); err != nil {
		return nil
	}

	return &lyrics
}
