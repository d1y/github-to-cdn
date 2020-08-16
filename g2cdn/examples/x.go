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

}
