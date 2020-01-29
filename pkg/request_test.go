package pkg_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	youtube "github.com/revett/youtube-comments/pkg"
	"github.com/stretchr/testify/require"
)

func TestDo(t *testing.T) {
	rp := loadRequestPages(t)

	calls := 0
	maxComments := 250
	token := "0e2346d3-158e-4713-921c-24350bf64532"
	videoID := "oS169nq8Prw"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query()
		require.Equal(t, token, qp.Get("key"))
		require.Equal(t, "100", qp.Get("maxResults"))
		require.Equal(t, "time", qp.Get("order"))
		require.Equal(t, "snippet,replies", qp.Get("part"))
		require.Equal(t, "oS169nq8Prw", qp.Get("videoId"))

		io.WriteString(w, rp[calls])
		calls++
	}))
	defer srv.Close()

	ctls, err := youtube.Do(token, videoID, maxComments, youtube.WithCustomEndpoint(srv.URL))

	require.NoError(t, err)
	require.Len(t, ctls, 3)
	require.GreaterOrEqual(t, ctls.Len(), maxComments)
}

func loadRequestPages(t *testing.T) []string {
	rp := []string{}
	for i := 0; i < 3; i++ {
		b := helperLoadBytes(t, fmt.Sprintf("page_%d.json", i))
		rp = append(rp, string(b))
	}
	return rp
}

func helperLoadBytes(t *testing.T, name string) []byte {
	p := filepath.Join("testdata", name)
	b, err := ioutil.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
