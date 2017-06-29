package ytcomments_test

import (
	"os"
	"testing"

	ytcomments "github.com/revett/go-yt-comments"
)

func TestAPI(t *testing.T) {
	t.Run("FetchComments()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			key := os.Getenv("API_KEY")
			maxComments := 250
			videoID := "oS169nq8Prw"

			api := ytcomments.NewAPI(key)

			commentThreadLists, err := api.FetchComments(videoID, maxComments)
			if err != nil {
				t.Fatalf(
					"Expected .FetchComments() to not return an error, got: '%s'", err,
				)
			}

			expectedLength := 3
			if len(commentThreadLists) != expectedLength {
				t.Fatalf(
					"Expected number of CommentThreadLists returned to be '%d', got: '%d'",
					expectedLength,
					len(commentThreadLists),
				)
			}
		})

		t.Run("SadCase", func(t *testing.T) {
			key := ""
			maxComments := 250
			videoID := "oS169nq8Prw"

			api := ytcomments.NewAPI(key)

			_, err := api.FetchComments(videoID, maxComments)
			if err == nil {
				t.Fatal("Expected .FetchComments() to return an error, got: 'nil'")
			}
		})
	})
}
