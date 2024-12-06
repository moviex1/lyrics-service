package request

type UpdateLyricsRequest struct {
	SongId  string `json:"songId"`
	Content string `json:"content"`
}
