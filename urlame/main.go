package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

var numberregex = regexp.MustCompile("^\\d+$")
var profilepageregex = regexp.MustCompile("^/(u|user|profile)/[^/]+/?$")
var titleregex = regexp.MustCompile("^[a-z0-9-]+$")

var hashregex = regexp.MustCompile("^[a-zA-Z0-9]+$")
var hashlens = []int{32, 40, 64, 128}

var exts = []string{".js", ".css", ".png", ".jpg", ".jpeg", ".svg", ".gif", ".mp3", ".mp4", ".rss", ".ttf", ".woff", ".woff2", ".eot", ".pdf"}
var paths = []string{"/static/", "/assets/", "/wp-content/", "/blog/", "/product/", "/docs/", "/support/"}

func main() {
	printNormalized := flag.Bool("n", false, "print the normalized version of the urls (for debugging)")
	flag.Usage = func() {
		fmt.Printf("cat urls.txt | %s [OPTIONS]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	seen := map[string]bool{}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		urlstr := stdin.Text()
		if u, err := url.Parse(urlstr); err == nil {
			if lamefiletype(u) || lamedir(u) || profilepage(u) {
				//skip those that we can be certain are lame
				continue
			}
			normalized := normalizeUrl(urlstr)
			if seen[normalized] {
				continue
			} else {
				seen[normalized] = true
			}
			if *printNormalized {
				fmt.Println(normalized)
			} else {
				fmt.Println(urlstr)
			}
		}
	}
}

func lamefiletype(u *url.URL) bool {
	filetype := strings.ToLower(path.Ext(u.Path))
	for _, ext := range exts {
		if filetype == ext {
			return true
		}
	}
	return false
}

func lamedir(u *url.URL) bool {
	for _, p := range paths {
		if strings.HasPrefix(u.Path, p) {
			return true
		}
	}
	return false
}

func profilepage(u *url.URL) bool {
	if profilepageregex.MatchString(u.Path) {
		return true
	}
	return false
}

func normalizeUrl(urlstr string) string {
	//this method needs the original string instead of a url as it may return it unchanged
	if u, err := url.Parse(urlstr); err == nil {
		newvals := url.Values{}
		for key := range u.Query() {
			newvals.Set(key, "%P%")
		}
		return newUrl(u, normalizePath(u.Path), newvals)
	}
	return urlstr
}

func normalizePath(path string) string {
	normalized := ""
	for _, part := range strings.Split(path, "/") {
		if strings.TrimSpace(part) == "" {
			continue
		}
		normalized += "/" + normalizeItem(part)
	}
	return normalized
}

func normalizeItem(item string) string {
	// it's unlikely that we have urls with %X% in them which we would miss here
	if numberregex.MatchString(item) {
		return "%N%"
	} else if postitle(item) {
		return "%%T%"
	} else if hash(item) {
		return "%H%"
	}
	return item
}

func postitle(str string) bool {
	if !titleregex.MatchString(str) {
		return false
	}
	if len(str) > 10 {
		return true
	}
	return strings.Count(str, "-") > 2
}

func hash(str string) bool {
	if hashregex.MatchString(str) {
		strlen := len(str)
		for _, l := range hashlens {
			if strlen == l {
				return true
			}
		}
	}
	return false
}

func newUrl(old *url.URL, path string, vals url.Values) string {
	return /*ignore scheme*/ cleanHostname(old) + path + "?" + vals.Encode() + "#" + old.Fragment
}

func cleanHostname(u *url.URL) string {
	if u.Port() == "80" || u.Port() == "443" {
		return u.Hostname()
	}
	return u.Host
}
