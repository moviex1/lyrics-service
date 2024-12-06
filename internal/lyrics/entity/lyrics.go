package entity

type Lyrics struct {
	Id        string `db:"id"`
	SongId    string `db:"song_id"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
