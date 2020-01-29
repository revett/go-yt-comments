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

## Credentials

You will need a YouTube API token, see the [docs](https://developers.google.com/youtube/v3/docs/).

## Options

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

## Context

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

## Status

[![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/revett/youtube-comments/Test/master?style=flat-square)](https://github.com/revett/youtube-comments/actions?query=workflow%3ATest)

## Other

If you would like more control over the requests being made, then have a read of the official
Google package:

```
https://godoc.org/google.golang.org/api/youtube/v3
```
