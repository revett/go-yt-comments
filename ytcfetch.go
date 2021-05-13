package ytcfetch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	defaultMaxComments = 500
	endpoint           = "https://www.googleapis.com/youtube/v3/commentThreads"
	maxResultsPerPage  = "100"
)

type (
	// Option is a functional option for configuring the Do function.
	Option func(*config)

	// HTTPClient is TODO
	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	config struct {
		httpClient  HTTPClient
		maxComments int
		token       string
		videoID     string
	}
)

// Do is TODO.
func Do(token string, videoID string, opts ...Option) (CommentThreadLists, error) {
	cfg := config{
		httpClient:  http.DefaultClient,
		maxComments: defaultMaxComments,
		token:       token,
		videoID:     videoID,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	req, err := formRequest(cfg, "")
	if err != nil {
		return nil, err
	}

	lists := CommentThreadLists{}
	return fetch(cfg, req, lists)
}

func WithHTTPClient(c HTTPClient) Option {
	return func(cfg *config) {
		cfg.httpClient = c
	}
}

func fetch(cfg config, req *http.Request, lists CommentThreadLists) (CommentThreadLists, error) {
	resp, err := cfg.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error when attempting to close response body: %s", err)
		}
	}()

	if resp.StatusCode > 200 {
		return nil, fmt.Errorf(
			"non-200 status code returned from youtube api, got: '%d'", resp.StatusCode,
		)
	}

	var list CommentThreadList
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	lists = append(lists, list)

	if len(list.Items) == 0 {
		return lists, nil
	}

	if list.NextPageToken == "" {
		return lists, nil
	}

	if lists.Len() >= cfg.maxComments {
		return lists, nil
	}

	nextPageReq, err := formRequest(cfg, list.NextPageToken)
	if err != nil {
		return nil, err
	}

	return fetch(cfg, nextPageReq, lists)
}

func formRequest(cfg config, nextPageToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	qp := req.URL.Query()

	qp.Add("key", cfg.token)
	qp.Add("maxResults", maxResultsPerPage)
	qp.Add("order", "time")
	qp.Add("part", "snippet,replies")
	qp.Add("videoId", cfg.videoID)

	if nextPageToken != "" {
		qp.Add("pageToken", nextPageToken)
	}

	req.URL.RawQuery = qp.Encode()
	return req, nil
}
