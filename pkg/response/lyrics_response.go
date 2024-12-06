package response

type LyricsResponse struct {
	Id        string `json:"id"`
	SongId    string `json:"songId"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
