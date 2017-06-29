package api

import "time"

type (
	// CommentSnippet is a struct representation of the YouTube model.
	CommentSnippet struct {
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
)
