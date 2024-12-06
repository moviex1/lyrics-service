package lyrics

import "example.com/internal/lyrics/entity"

func (repo *lyricsRepository) FindBySongId(songId string) *entity.Lyrics {
	var lyrics entity.Lyrics

	if err := repo.conn.Get(&lyrics, "SELECT id, song_id, content, created_at, updated_at FROM lyrics WHERE song_id = $1", songId); err != nil {
		return nil
	}

	return &lyrics
}
