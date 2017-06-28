package ytcomments

import (
	"encoding/json"
	"net/http"

	"github.com/revett/go-yt-comments/models/api"
)

type (
	// API is the top level struct used to interact with this package, it requires
	// an API key.
	API struct {
		key string
	}
)

const (
	baseURI    = "https://www.googleapis.com/youtube/v3/commentThreads"
	maxResults = "100"
)

// NewAPI receives an API key as an argument, and returns a new API struct.
func NewAPI(key string) API {
	return API{
		key: key,
	}
}

// FetchComments receives a YouTube Video ID as an arguement, and returns a
// slice of Comment structs.
func (a API) FetchComments(videoID string) (*api.CommentThreadList, error) {
	client := http.Client{}

	request, err := a.formRequest(videoID)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var commentThreadList api.CommentThreadList

	err = json.NewDecoder(response.Body).Decode(&commentThreadList)
	if err != nil {
		return nil, err
	}

	return &commentThreadList, nil
}

func (a API) formRequest(videoID string) (*http.Request, error) {
	request, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return nil, err
	}

	queryParams := request.URL.Query()

	queryParams.Add("key", a.key)
	queryParams.Add("maxResults", maxResults)
	queryParams.Add("order", "time")
	queryParams.Add("part", "snippet,replies")
	queryParams.Add("videoId", videoID)

	request.URL.RawQuery = queryParams.Encode()

	return request, nil
}
