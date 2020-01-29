<p align="center">
  <img src="https://images.unsplash.com/photo-1497015289639-54688650d173?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=3889&q=80" width="100%">
</p>

<h1 align="center">
  youtube-comments
</h1>

<p align="center">
  Go package for quickly fetching video comments from the YouTube API
</p>

## Usage

The example below fetches a maxium of 250 comments for the this
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

## Credentials

You will need a YouTube API key, see the [docs](https://developers.google.com/youtube/v3/docs/).
