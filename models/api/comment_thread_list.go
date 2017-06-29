package api

type (
	// CommentThreadList is a struct representation of the YouTube model.
	CommentThreadList struct {
		Kind          string          `json:"kind"`
		ETAG          string          `json:"etag"`
		NextPageToken string          `json:"nextPageToken"`
		PageInfo      PageInfo        `json:"pageInfo"`
		Items         []CommentThread `json:"items"`
	}
)
