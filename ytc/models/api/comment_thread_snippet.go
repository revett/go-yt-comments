package api

type (
	// CommentThreadSnippet is a struct representation of the YouTube model.
	CommentThreadSnippet struct {
		ChannelID       string  `json:"channelId"`
		VideoID         string  `json:"videoId"`
		TopLevelComment Comment `json:"topLevelComment"`
		CanReply        bool    `json:"canReply"`
		TotalReplyCount int64   `json:"totalReplyCount"`
		IsPublic        bool    `json:"isPublic"`
	}
)
