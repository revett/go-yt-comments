# ytcfetch [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/revett/ytcfetch) [![Go Report Card](https://goreportcard.com/badge/github.com/revett/ytcfetch)](https://goreportcard.com/report/github.com/revett/ytcfetch)

📹 `ytcfetch` - **YouTube Comment Fetcher** is a small convenience Go package for
quickly fetching the comments for a specific YouTube video. It has zero
dependencies.

## Usage

View [**full documentation** (godoc) →](https://pkg.go.dev/github.com/revett/ytcfetch)

### Example

The example below fetches a maxium of 250 comments for this
[video](https://www.youtube.com/watch?v=oS169nq8Prw).

```golang
package main

import (
  "context"
  "log"

  youtube "github.com/revett/youtube-comments/pkg"
)

func main() {
  ctx := context.Background()
  ctls, err := youtube.Do(ctx, "TOKEN", "oS169nq8Prw", 250)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("got %d comments", ctls.Len())
}
```

### Credentials

You will need a YouTube API token, see the [docs](https://developers.google.com/youtube/v3/docs/).

### Options

`youtube.Do()` takes zero or more `opts` that allow for the modification for the
underlying API client.

The following `opt` overrides the default API endpoint used:

```golang
func main() {
  ctx := context.Background()
  e := "https://example.com"
  _, err := youtube.Do(ctx, "TOKEN", "oS169nq8Prw", 250, youtube.WithCustomEndpoint(e))
}
```

### Context

`youtube.Do()` takes a `context.Context`, which allows you to cancel the request, or specify a
sensible timeout etc.

```golang
func main() {
  ctx := context.Background()
  ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
  defer cancel()

  _, err := youtube.Do(ctx, "TOKEN", "oS169nq8Prw", 250)
}
```

## Examples

- TODO

## Dependencies

- None

## Other

If you would like more control over the requests being made, then have a read of the official
Google package:

```
https://godoc.org/google.golang.org/api/youtube/v3
```

## Credit

- TODO
