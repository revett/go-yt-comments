package pkg

import "time"

// CommentSnippetAuthorChannelID is a struct representation of the YouTube
// model.
type CommentSnippetAuthorChannelID struct {
	Value string `json:"value"`
}

// CommentSnippet is a struct representation of the YouTube model.
type CommentSnippet struct {
	AuthorDisplayName     string                        `json:"authorDisplayName"`
	AuthorProfileImageURL string                        `json:"authorProfileImageUrl"`
	AuthorChannelURL      string                        `json:"authorChannelUrl"`
	AuthorChannelID       CommentSnippetAuthorChannelID `json:"authorChannelId"`
	ChannelID             string                        `json:"channelId"`
	VideoID               string                        `json:"videoId"`
	TextDisplay           string                        `json:"textDisplay"`
	TextOriginal          string                        `json:"textOriginal"`
	ParentID              string                        `json:"parentId"`
	CanRate               bool                          `json:"canRate"`
	ViewerRating          string                        `json:"viewerRating"`
	LikeCount             int64                         `json:"likeCount"`
	ModerationStatus      string                        `json:"moderationStatus"`
	PublishedAt           time.Time                     `json:"publishedAt"`
	UpdatedAt             time.Time                     `json:"updatedAt"`
}

// CommentThreadList is a struct representation of the YouTube model.
type CommentThreadList struct {
	Kind          string          `json:"kind"`
	ETAG          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	PageInfo      PageInfo        `json:"pageInfo"`
	Items         []CommentThread `json:"items"`
}

// CommentThreadLists allows for helper funcs on a slice of CommentThreadList structs.
type CommentThreadLists []CommentThreadList

// Len returns the number of comments contained in the slice of CommentThreadList structs.
func (c CommentThreadLists) Len() int {
	i := 0
	for _, ctl := range c {
		for _, ct := range ctl.Items {
			i++
			i += len(ct.Replies.Comments)
		}
	}
	return i
}

type (
	// CommentThreadReplies is a struct representation of the YouTube model.
	CommentThreadReplies struct {
		Comments []Comment `json:"comments"`
	}
)

// CommentThreadSnippet is a struct representation of the YouTube model.
type CommentThreadSnippet struct {
	ChannelID       string  `json:"channelId"`
	VideoID         string  `json:"videoId"`
	TopLevelComment Comment `json:"topLevelComment"`
	CanReply        bool    `json:"canReply"`
	TotalReplyCount int64   `json:"totalReplyCount"`
	IsPublic        bool    `json:"isPublic"`
}

// CommentThread is a struct representation of the YouTube model.
type CommentThread struct {
	Kind    string               `json:"kind"`
	ETAG    string               `json:"etag"`
	ID      string               `json:"id"`
	Snippet CommentThreadSnippet `json:"snippet"`
	Replies CommentThreadReplies `json:"replies"`
}

// Comment is a struct representation of the YouTube model.
type Comment struct {
	Kind    string         `json:"kind"`
	ETAG    string         `json:"etag"`
	ID      string         `json:"id"`
	Snippet CommentSnippet `json:"snippet"`
}

// PageInfo is a struct representation of the YouTube model.
type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
