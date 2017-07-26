package api

type (
	// CommentThread is a struct representation of the YouTube model.
	CommentThread struct {
		Kind    string               `json:"kind"`
		ETAG    string               `json:"etag"`
		ID      string               `json:"id"`
		Snippet CommentThreadSnippet `json:"snippet"`
		Replies CommentThreadReplies `json:"replies"`
	}
)
