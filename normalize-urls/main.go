package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	unique := flag.Bool("u", false, "unique")
	donthttps := flag.Bool("h", false, "do not set all urls to https://")
	flag.Parse()
	stdin := bufio.NewScanner(os.Stdin)
	seen := map[string]bool{}
	for stdin.Scan() {
		if u, err := url.Parse(stdin.Text()); err == nil {
			scheme := u.Scheme
			if !*donthttps {
				scheme = "https"
			}
			cleaned := scheme + "://" + cleanHostname(u) + cleanPath(u) + cleanQuery(u) + cleanFragment(u)
			if *unique {
				if seen[cleaned] {
					continue
				}
				seen[cleaned] = true
			}
			fmt.Println(cleaned)
		}
	}
}

func cleanHostname(u *url.URL) string {
	if u.Port() == "80" || u.Port() == "443" {
		return u.Hostname()
	}
	return u.Host
}

func cleanPath(u *url.URL) string {
	if u.EscapedPath() == "/" {
		return ""
	}
	return u.EscapedPath()
}

func cleanQuery(u *url.URL) string {
	if u.RawQuery == "" {
		return ""
	}
	return "?" + u.RawQuery
}

func cleanFragment(u *url.URL) string {
	if u.RawFragment == "" {
		return ""
	}
	return "#" + u.RawFragment
}
