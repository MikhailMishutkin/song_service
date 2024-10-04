package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"song_service/internal/models"
	"time"
)

func (h *HTTPSongHandle) GetSongInfo(input *models.SongInput) (song *models.Song, err error) {

	params := url.Values{}
	params.Add("group", input.Group)
	params.Add("song", input.Song)

	url := "https://localhost:9000/info?" + params.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		nil,
	)
	FailOnErrors(err, "error with NewRequestWithContext")

	response, err := http.DefaultClient.Do(req)
	FailOnErrors(err, "error when executing the request to API")

	content, err := io.ReadAll(response.Body)
	FailOnErrors(err, "fail to read response")

	err = json.Unmarshal(content, &song)
	FailOnErrors(err, "corrupt json data")

	return song, nil
}
