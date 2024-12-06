package request

type CreateLyricsRequest struct {
	SongId  string `json:"songId" validate:"required,uuid"`
	Content string `json:"content" validate:"required"`
}
