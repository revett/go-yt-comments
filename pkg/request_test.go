package pkg_test

import (
	"os"
	"testing"

	youtube "github.com/revett/youtube-comments/pkg"
	"github.com/stretchr/testify/require"
)

func TestDo(t *testing.T) {
	token := os.Getenv("TOKEN")
	maxComments := 250
	videoID := "oS169nq8Prw"

	ctls, err := youtube.Do(token, videoID, maxComments)

	require.NoError(t, err)
	require.Len(t, ctls, 3)
}
