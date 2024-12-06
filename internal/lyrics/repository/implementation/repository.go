package lyrics

import (
	postgres "example.com/db"
	"example.com/internal/lyrics/repository"
	"github.com/jmoiron/sqlx"
)

var _ repository.LyricsRepository = (*lyricsRepository)(nil)

type lyricsRepository struct {
	conn *sqlx.DB
}

func NewLyricsRepository() *lyricsRepository {
	return &lyricsRepository{
		conn: postgres.GetConnection(),
	}
}
