package ytc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/revett/go-yt-comments/ytc/models/api"
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

// FetchComments receives a YouTube Video ID and a MaxComments integer as
// arguments. The Video ID is used to know what video to fetch comments for.
// The MaxComments integer is used to know how many comments to fetch. The
// function returns an array of CommentThreadList YouTube structs.
func (a API) FetchComments(videoID string, maxComments int) ([]api.CommentThreadList, error) {
	request, err := a.formRequest(videoID, "")
	if err != nil {
		return nil, err
	}

	var commentThreadLists []api.CommentThreadList

	return a.fetch(request, maxComments, commentThreadLists)
}

func (a API) countComments(commmentThreadLists []api.CommentThreadList) int {
	count := 0

	for _, commmentThreadList := range commmentThreadLists {
		for _, commentThread := range commmentThreadList.Items {
			count++
			count += len(commentThread.Replies.Comments)
		}
	}

	return count
}

func (a API) fetch(request *http.Request, maxComments int, commmentThreadLists []api.CommentThreadList) ([]api.CommentThreadList, error) {
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode > 200 {
		return nil, fmt.Errorf(
			"Non-200 status code returned from YouTube API, got: '%d'",
			response.StatusCode,
		)
	}

	var commentThreadList api.CommentThreadList

	err = json.NewDecoder(response.Body).Decode(&commentThreadList)
	if err != nil {
		return nil, err
	}

	commmentThreadLists = append(commmentThreadLists, commentThreadList)

	if len(commentThreadList.Items) == 0 {
		return commmentThreadLists, nil
	}

	if commentThreadList.NextPageToken == "" {
		return commmentThreadLists, nil
	}

	commentCount := a.countComments(commmentThreadLists)

	if commentCount < maxComments {
		nextRequest, err := a.formRequest(
			commentThreadList.Items[0].Snippet.VideoID,
			commentThreadList.NextPageToken,
		)
		if err != nil {
			return nil, err
		}

		return a.fetch(nextRequest, maxComments, commmentThreadLists)
	}

	return commmentThreadLists, nil
}

func (a API) formRequest(videoID string, nextPageToken string) (*http.Request, error) {
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

	if nextPageToken != "" {
		queryParams.Add("pageToken", nextPageToken)
	}

	request.URL.RawQuery = queryParams.Encode()

	return request, nil
}
