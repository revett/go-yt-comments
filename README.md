# go-yt-comments

![GitHub tag](https://img.shields.io/github/tag/revett/go-yt-comments.svg?style=flat)
[![CircleCI](https://circleci.com/gh/revett/go-yt-comments/tree/master.svg?style=shield)](https://circleci.com/gh/revett/go-yt-comments/tree/master)
[![GoReportCard](https://goreportcard.com/badge/github.com/revett/go-yt-comments)](https://goreportcard.com/report/github.com/revett/go-yt-comments)

Tiny Golang package for retrieving video comments from the YouTube API.

## API Access

You will need a YouTube API key, see the [docs](https://developers.google.com/youtube/v3/docs/).

## Usage

```golang
package main

import (
  "github.com/revett/go-yt-comments/ytc"
)

func main() {
  key := os.Getenv("API_KEY")
  maxComments := 250
  videoID := "oS169nq8Prw"

  api := ytc.NewAPI(key)

  commentThreadLists, err := api.FetchComments(videoID, maxComments)
  if err != nil {
    panic(err)
  }

  fmt.Println(commentThreadLists)
}
```