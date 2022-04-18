package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var wwwregex = regexp.MustCompile("^www\\.")

func main() {
	routines := flag.Uint("r", 10, "go routines (concurrency)")
	cookies := flag.String("c", "", "cookies e.g. session=abc123")
	authheader := flag.String("a", "", "Authorization header e.g. Bearer: abc123")
	flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	urls := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < int(*routines); i++ {
		wg.Add(1)
		c := getClient()
		go func() {
			defer wg.Done()
			for u := range urls {
				if req, err := http.NewRequest("GET", u, nil); err == nil {
					if *cookies != "" {
						req.Header.Set("Cookie", *cookies)
					}
					if *authheader != "" {
						req.Header.Set("Authorization", *authheader)
					}
					if resp, err := c.Do(req); err == nil {
						if loc, err := resp.Location(); err == nil {
							if uu, err := url.Parse(u); err == nil {
								if isInterestingRedirect(uu, loc) {
									fmt.Println("from:\t" + u)
									fmt.Println("to:\t" + loc.String())
									fmt.Println()
								}
							}
						}
					}
				}

			}
		}()
	}
	for s.Scan() {
		urls <- s.Text()
	}
	close(urls)
}

func isInterestingRedirect(f, t *url.URL) bool {
	fromhost := normalHost(f)
	tohost := normalHost(t)
	from := fromhost + normalPath(f) + normalQuery(f) + normalFragment(f)
	to := tohost + normalPath(t) + normalQuery(t) + normalFragment(t)
	if from == to {
		//seems to be http to https or normalization
		return false
	}
	home := []string{"https://" + fromhost, "https://" + fromhost + "/"}
	for _, loc := range home {
		if loc == to {
			// redirects to home are like 404
			return false
		}
	}
	if f1, err := url.QueryUnescape(f.String()); err == nil {
		if f2, err := url.QueryUnescape(t.String()); err == nil {
			if f1 == f2 {
				//ignore % normalization
				return false
			}
		}
	}
	return true
}

func normalHost(u *url.URL) string {
	return strings.ToLower(wwwregex.ReplaceAllString(u.Hostname(), ""))
}

func normalPath(u *url.URL) string {
	path := u.EscapedPath()
	if path == "" {
		return "/"
	}
	return path
}

func normalQuery(u *url.URL) string {
	query := u.Query().Encode()
	if query == "" {
		return ""
	}
	return "?" + query
}

func normalFragment(u *url.URL) string {
	if u.Fragment == "" {
		return ""
	}
	return "#" + u.Fragment
}

func getClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       time.Second * 10,
	}
}
