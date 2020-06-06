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
