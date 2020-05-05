/*
** link: https://www.jsdelivr.com/?docs=gh
*/

const testLink = link=> {
  const regex = /^(((ht|f)tps?):\/\/)?[\w-]+(\.[\w-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$/
  return regex.test(link)
}

const isObject = x => {
  return typeof x === 'object' && x !== null
}

const isGithubRepo = repo => {
  const regex = /(https|git@[-\w.]+):(\/\/)?(.*?)(\/?|[-\d\w._]+?)$/
  const isRepo = /^(?!.*https|git).*\w+(\/)\w+.*$/
  return regex.test(repo) || isRepo.test(repo)
}

const githubSiteTest = link=> {
  const x = /(github\.(io|com)|raw\.githubusercontent\.com)/
  return [x.test(link), x]
}

// https://d1y. => d1y
const formatIoUserName = link=> {
  const l = /(http|https):\/\//.exec(link)
  const p = l[1]
  const lenZero = link.length - 1
  let result
  if (p === 'http') {
    result = link.substring(7, lenZero)
  } else if (p === 'https') {
    result = link.substring(8, lenZero)
  }
  return result
}

const isLink = link=> {

  // https://iptv-org.github.io/iptv/index.m3u
  // https://github.com/video-dev/hls.js/blob/master/src/empty.js
  // https://raw.githubusercontent.com/bukinoshita/is-github-repo/master/package.json

  if (!testLink(link)) return false
  const [ isGithubLink, regex ] = githubSiteTest(link)
  if (isGithubLink) {
    const result = {
      username: '',
      repo: '',
      branch: 'master',
      path: ''
    }
    const pm = regex.exec(link)
    const site = pm[0]
    const index = pm['index']
    const siteLen = site.length
    const data = link.substring(index + siteLen)
    const sp = data.substring(1).split('/')
    if (site === 'github.com' || site === 'raw.githubusercontent.com') {
      result['username'] = sp[0]
      result['repo'] = sp[1]
    }
    if (site === 'github.io')  {
      // https://d1y. => d1y
      const user = formatIoUserName(link.substring(0, index))
      result['username'] = user
      const linkAddOneLen = index + siteLen
      const seg = link.substr(linkAddOneLen, 1)
      if (seg === '/') {
        // iptv/index.m3u => iptv
        const info = link.substring(linkAddOneLen + 1)
        const split = info.split('/')
        const repo = split[0]
        result['repo'] = repo
        result['branch'] = 'gh-pages'
        // iptv/index.m3u => index.m3u
        let path = info.substring(repo.length + 1)
        if (path[0] === '/') path = path.substring(1)
        result['path'] = path
        return result
      }
    } else if (site === 'github.com') {
      if (data[0] === '/') {
        if (sp[2] === 'blob') {
          result['branch'] = sp[3]
          result['path'] = sp.filter((item, index)=> {
            return index > 3
          }).join('/')
          return result
        }
      }
    } else if (site === 'raw.githubusercontent.com') {
      if (data[0] === '/') {
        result['branch'] = sp[2]
        result['path'] = sp.filter((item, index)=> {
          return index > 2
        }).join('/')
        return result
      }
    }
  }
  return false
}

const toCDNLink = link=> {
  let result = {}
  if (isObject(link)) {
    if (link.hasOwnProperty('username')) {
      result['username'] = link['username']
    } else {
      result['username'] = 'd1y'
    }
    if (link.hasOwnProperty('repo')) {
      result['repo'] = link['repo']
    } else {
      result['repo'] = 'ass-we-can'
    }
    if(link.hasOwnProperty('path')) {
      result['path'] = link['path']
    } else {
      result['path'] = 'src/lite.json'
    }
    if (link.hasOwnProperty('branch')) {
      result['branch'] = link['branch']
    }
  } else if (typeof link === 'string') {
    const data = isLink(link)
    if (!data) return false
    result = data
  } else return false
  let s =`https://cdn.jsdelivr.net/gh/${ result.username }/${ result.repo }`
  const b = `@${ result.branch }/`
  s += `${ result.branch ? b : '/' }${ result.path }`
  return s
}

module.exports = toCDNLink