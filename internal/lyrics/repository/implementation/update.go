package lyrics

import (
	"time"

	"example.com/internal/lyrics/entity"
	"example.com/pkg/constants"
)

func (repo *lyricsRepository) Update(lyrics *entity.Lyrics) {
	repo.conn.MustExec("UPDATE lyrics SET content = $1, updated_at = $2 WHERE id = $3", lyrics.Content, time.Now().Format(constants.DateFormat), lyrics.Id)
}
