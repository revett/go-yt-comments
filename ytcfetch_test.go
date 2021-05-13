package ytcfetch_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/revett/ytcfetch"
	"github.com/stretchr/testify/require"
)

func TestDo(t *testing.T) {
	t.Parallel()

	token := "0e2346d3-158e-4713-921c-24350bf64532"
	videoID := "oS169nq8Prw"

	tests := map[string]struct {
		client ytcfetch.HTTPClient
		want   int
		err    require.ErrorAssertionFunc
	}{
		"Simple": {
			client: mockClient{},
			err:    require.NoError,
			want:   13,
		},
	}

	for n, testCase := range tests {
		tc := testCase

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			lists, err := ytcfetch.Do(
				token, videoID, ytcfetch.WithHTTPClient(tc.client),
			)

			tc.err(t, err)
			require.Equal(t, tc.want, lists.Len())
		})
	}
}

type mockClient struct{}

func (m mockClient) Do(req *http.Request) (*http.Response, error) {
	b := helperLoadBytes(fmt.Sprintf("page_%d.json", 0))
	r := ioutil.NopCloser(bytes.NewReader([]byte(b)))

	return &http.Response{
		Body: r,
	}, nil
}

func helperLoadBytes(name string) []byte {
	p := filepath.Join("testdata", name)

	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
