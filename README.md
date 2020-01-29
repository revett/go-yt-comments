<p align="center">
  <img src="./docs/header.png" width="100%">
</p>

<h1 align="center">
  youtube-comments
</h1>

<p align="center">
  Go package for quickly fetching video comments from the YouTube API
</p>

## Usage

The example below fetches a maxium of 250 comments for this
[video](https://www.youtube.com/watch?v=oS169nq8Prw).

```golang
package main

import (
  "log"

  youtube "github.com/revett/youtube-comments/pkg"
)

func main() {
  ctls, err := youtube.Do("TOKEN", "oS169nq8Prw", 250)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("got %d comments", ctls.Len())
}
```

## Options

`youtube.Do()` takes zero or more `opts` that allow for the modification for the
underlying API client.

The following `opt` overrides the default API endpoint used:

```golang
func main() {
  e := "https://example.com"
  _, err := youtube.Do("TOKEN", "oS169nq8Prw", 250, youtube.WithCustomEndpoint(e))
}
```

## Credentials

You will need a YouTube API token, see the [docs](https://developers.google.com/youtube/v3/docs/).

## Other

If you would like more control over the requests being made, then have a read of the official
Google package:

```
https://godoc.org/google.golang.org/api/youtube/v3
```
