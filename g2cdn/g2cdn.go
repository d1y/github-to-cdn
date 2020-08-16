// create by d1y<chenhonzhou@gmail.com>

package g2cdn

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	// DefaultBranch 默认分支
	DefaultBranch = `master`
)

// Config 配置
type Config struct {
	Username string // 用户名
	Repo     string // 仓库
	Path     string // 路径
	Branch   string // 分支
}

// Face test
type Face struct {
	Config Config // 配置
	// ToString func() string // 转换成字符串
}

// New 创建一个配置
func New(config Config) (Face, error) {
	var username, repo, path, branch = config.Username, config.Repo, config.Path, config.Branch
	if isEmpty(username) || isEmpty(repo) || isEmpty(path) {
		return Face{}, errors.New("config error")
	}
	if isEmpty(branch) {
		config.Branch = DefaultBranch
	}
	return Face{
		Config: config,
	}, nil
}

// ToString 转换成 `url`
func (face Face) ToString() string {
	var config = face.Config
	var username, repo, path, branch = config.Username, config.Repo, config.Path, config.Branch
	if len(path) == 0 {
		fmt.Println("path: ", path)
		return ""
	}
	if path[0] == '/' {
		path = path[1:]
	}
	var result = fmt.Sprintf(`https://cdn.jsdelivr.net/gh/%v/%v@%v/%v`, username, repo, branch, path)
	return result
}

// DomainType 域名类型
type DomainType int

const (
	// GithubIO github.io
	GithubIO DomainType = 0
	// GithubCOM github.com
	GithubCOM DomainType = 1
	// GithubRAW raw.githubusercontent.com
	GithubRAW DomainType = 2
)

// 包含的域名
var domains = []string{
	"github.io",
	"github.com",
	"raw.githubusercontent.com",
}

// Parse 解析 `url`
// [hasBlob] 当你的目录中存在 /user/repo/[blob]/branch, 加的兼容处理
func Parse(rawURL string, hasBlob bool) (Face, error) {
	U, err := url.Parse(rawURL)
	if err != nil {
		return Face{}, err
	}
	var hostname = U.Hostname()
	var d int = -1
	for index, item := range domains {
		if item == hostname {
			d = index
		} else {
			var dd = strings.Split(hostname, ".")
			if len(dd) == 3 && dd[1] == "github" && dd[2] == "io" {
				d = 0
			}
		}
	}
	if d == -1 {
		return Face{}, errors.New("match domain is error")
	}
	// fmt.Println("d", d)
	var conf Config
	var Xpath = U.Path
	var Xpaths = strings.Split(Xpath, "/")
	// var x = len(Xpaths)
	if d >= 1 {
		conf.Username = Xpaths[1]
		conf.Repo = Xpaths[2]
		var offset = 3
		if hasBlob {
			offset = 4
		}
		conf.Branch = Xpaths[offset]
		pathSp := Xpaths[offset+1:]
		vpath := strings.Join(pathSp, "/")
		// fmt.Println("vpath", vpath)
		conf.Path = vpath
	} else {
		var hh = strings.Split(hostname, ".")
		if len(hh) >= 1 {
			conf.Username = hh[0]
		}
		conf.Repo = Xpaths[1]
		p := strings.Join(Xpaths[2:], "/")
		conf.Branch = DefaultBranch
		conf.Path = p
	}
	pp := Face{
		Config: conf,
	}
	// var ul = pp.ToString()
	// fmt.Println("ul", ul)
	return pp, err
}

// Easy 最简单的用法, 直接返回 `url` 如果错误返回空字符串
func Easy(u string, hasBlob bool) string {
	face, e := Parse(u, hasBlob)
	if e != nil {
		return ""
	}
	return face.ToString()
}

// Check 检测链接是否符合`github-domain`标椎
func Check(u string) bool {
	x, e := url.Parse(u)
	if e != nil {
		return false
	}
	h := x.Hostname()
	for _, item := range domains {
		if item == h {
			return true
		}
	}
	var arr = strings.Split(h, ".")
	return arr[1] == "github" && arr[2] == "io"
}

// 判断是否为空
func isEmpty(s string) bool {
	return len(s) == 0
}
