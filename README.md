你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
<p align="center">
  <img src="./docs/header.png" width="100%">
</p>

<h1 align="center">
  youtube-comments
</h1>

<p align="center">
  Go package for quickly fetching YouTube video comments
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
