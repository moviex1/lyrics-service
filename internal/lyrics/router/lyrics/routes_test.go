package lyrics_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "example.com/internal/lyrics/api/rest"
	"example.com/internal/lyrics/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateSuccess(t *testing.T) {
	service := new(LyricsServiceMock)
	service.On("UpdateLyrics", mock.AnythingOfType("*dto.UpdateLyricsDto"), "12345").Return(&dto.LyricsDto{
		Id:        "some-uuid",
		SongId:    "some-uuid",
		Content:   "content",
		CreatedAt: "some-data",
		UpdatedAt: "some-data",
	})

	controller := v1.NewLyricsController(service)
	gin.SetMode(gin.TestMode)

	updateLyricsDto := &dto.UpdateLyricsDto{
		Content: "Updated lyrics content",
	}
	body, err := json.Marshal(updateLyricsDto)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/lyrics/12345", bytes.NewBuffer(body))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(w)
	context.Request = req
	context.Params = gin.Params{{Key: "songId", Value: "12345"}}

	controller.UpdateLyrics(context)

	assert.Equal(t, http.StatusOK, w.Code)

	service.AssertCalled(t, "UpdateLyrics", mock.AnythingOfType("*dto.UpdateLyricsDto"), "12345")
}

type LyricsServiceMock struct {
	mock.Mock
}

func (s *LyricsServiceMock) UpdateLyrics(updateDto *dto.UpdateLyricsDto, songId string) *dto.LyricsDto {
	args := s.Called(updateDto, songId)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.LyricsDto)
	}
	return nil
}

func (s *LyricsServiceMock) AddLyrics(dto *dto.CreateLyricsDto) *dto.LyricsDto {
	return nil
}

func (s *LyricsServiceMock) GetLyricsBySongId(songId string) *dto.LyricsDto {
	return nil
}

func (s *LyricsServiceMock) GetAllLyrics() []*dto.LyricsDto {
	return []*dto.LyricsDto{}
}
