package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/revett/youtube-comments/internal"
)

// Do receives a YouTube Video ID and a MaxComments integer as arguments. The Video ID is used to
// know what video to fetch comments for. The MaxComments integer is used to know how many comments
// to fetch. The function returns an array of CommentThreadList YouTube structs.
func Do(token string, videoID string, maxComments int, opts ...internal.ClientOption) (CommentThreadLists, error) {
	c := internal.NewClient(token, opts...)

	r, err := formRequest(c, videoID, "")
	if err != nil {
		return nil, err
	}

	var ctls CommentThreadLists
	return fetch(0, c, r, maxComments, ctls)
}

// WithCustomEndpoint specifies a different underlying API endpoint to use when making requests.
func WithCustomEndpoint(e string) internal.ClientOption {
	return func(c *internal.Client) {
		c.Endpoint = e
	}
}

func fetch(count int, c *internal.Client, req *http.Request, maxComments int, ctls CommentThreadLists) (CommentThreadLists, error) {
	count++

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		return nil, fmt.Errorf(
			"non-200 status code returned from youtube api, got: '%d'", resp.StatusCode,
		)
	}

	// // new

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// buffer := new(bytes.Buffer)
	// if err := json.Compact(buffer, body); err != nil {
	// 	return nil, err
	// }

	// if count == 4 {
	// 	fmt.Println(buffer)
	// }

	// var ctl CommentThreadList
	// err = json.Unmarshal(body, &ctl)
	// if err != nil {
	// 	return nil, err
	// }

	// // end new

	var ctl CommentThreadList
	err = json.NewDecoder(resp.Body).Decode(&ctl)
	if err != nil {
		return nil, err
	}

	ctls = append(ctls, ctl)

	if len(ctl.Items) == 0 || ctl.NextPageToken == "" {
		return ctls, nil
	}

	if ctls.Len() < maxComments {
		nextReq, err := formRequest(c, ctl.Items[0].Snippet.VideoID, ctl.NextPageToken)
		if err != nil {
			return nil, err
		}

		return fetch(count, c, nextReq, maxComments, ctls)
	}

	return ctls, nil
}

func formRequest(c *internal.Client, videoID string, nextPageToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	qp := req.URL.Query()

	qp.Add("key", c.Token)
	qp.Add("maxResults", internal.MaxResults)
	qp.Add("order", "time")
	qp.Add("part", "snippet,replies")
	qp.Add("videoId", videoID)

	if nextPageToken != "" {
		qp.Add("pageToken", nextPageToken)
	}

	req.URL.RawQuery = qp.Encode()
	return req, nil
}
