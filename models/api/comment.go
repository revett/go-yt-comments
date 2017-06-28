package api

type (
	// Comment is a struct representation of the YouTube model.
	Comment struct {
		Kind string `json:"kind"`
		ETAG string `json:"etag"`
		ID   string `json:"id"`
	}
)
