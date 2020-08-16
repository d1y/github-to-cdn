# github-to-cdn

Mainly convert github address to cdn address to speed up the website. cdn by [jsdelivr.com](https://www.jsdelivr.com)

![](https://www.jsdelivr.com/img/logo-horizontal.svg)

Support site types

- https://iptv-org.github.io/iptv/index.m3u
- https://github.com/video-dev/hls.js/blob/master/src/empty.js
- https://raw.githubusercontent.com/bukinoshita/is-github-repo/master/package.json

exmaples

```js
import ghCDN from 'github-to-cdn'

// the reutrn `string` type

const linkToObject = ghCDN({
  username: 'd1y',
  repo: 'ass-we-can',
  path: 'src/lite.json',
  branch: 'master'
})

const linkToString = ghCDN(`https://github.com/d1y/ass-we-can/blob/master/src/lite.json`)

console.log(linkToObject)
console.log(linkToString)
```

# g2cdn

```
go get -u github.com/d1y/github-to-cdn/g2cdn
```

the examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/d1y/github-to-cdn/g2cdn"
)

func main() {

	var config = g2cdn.Config{
		Username: "d1y",
		Repo:     "yoxi_data",
		Path:     "result/x.json",
	}
	t, err := g2cdn.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	// https://cdn.jsdelivr.net/gh/d1y/yoxi_data@master/result/x.json
	t.ToString()

	g2cdn.Parse(`https://github.com/d1y/github-to-cdn/blob/master/index.js`, true)

	ctx, _ := g2cdn.Parse(`https://iptv-org.github.io/iptv/index.m3u`, false)
	ctx.Config.Branch = "gh_pages"
	var u = ctx.ToString()
	fmt.Println("u", u)

	raw, _ := g2cdn.Parse(`https://raw.githubusercontent.com/bukinoshita/is-github-repo/master/package.json`, false)
	var r = raw.ToString()
	fmt.Println("r", r)

	var url = g2cdn.Easy(`https://raw.githubusercontent.com/bukinoshita/is-github-repo/master/package.json`, false)
	fmt.Println("url", url)

	var f = g2cdn.Check(`https://github.com/d1y/github-to-cdn/blob/master/index.js`)
	fmt.Println("f", f)

}

```