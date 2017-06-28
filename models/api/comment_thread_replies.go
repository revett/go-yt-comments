package api

type (
	// CommentThreadReplies is a struct representation of the YouTube model.
	CommentThreadReplies struct {
		Comments []Comment `json:"comments"`
	}
)
