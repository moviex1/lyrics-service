package dto

type LyricsDto struct {
	Id        string
	SongId    string
	Content   string
	CreatedAt string
	UpdatedAt string
}

type CreateLyricsDto struct {
	SongId  string
	Content string
}

type UpdateLyricsDto struct {
	Content string
}
