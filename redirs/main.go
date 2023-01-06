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
	routines := flag.Uint("r", 1, "go routines (concurrency)")
	cookies := flag.String("c", "", "cookies e.g. session=abc123")
	authheader := flag.String("a", "", "Authorization header e.g. Bearer: abc123")
	typical := flag.Bool("typical", false, "only check URLs that look like typical redirects")
	interesting := flag.Bool("interesting", false, "only output interesting URLs")
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
				fmt.Print(".")
				if *typical && !looksLikeRedirect(u) {
					continue
				}
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
								if !isLameRedirect(uu, loc) || (*interesting && isInterestingRedirect(uu, loc)) {
									if commonsubstr := getMaxLengthCommonString(u, loc.String()); len(commonsubstr) < 3 {
										continue
									} else {
										if !*interesting && isInterestingRedirect(uu, loc) {
											fmt.Println("INTERESTING REDIRECT!")
										}
										fmt.Printf("from:\t%s\nto:\t%s\n->\t%s\n\n", u, loc.String(), commonsubstr)
									}
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

func isLameRedirect(f, t *url.URL) bool {
	fromhost := normalHost(f)
	tohost := normalHost(t)
	from := fromhost + normalPath(f) + normalQuery(f) + normalFragment(f)
	to := tohost + normalPath(t) + normalQuery(t) + normalFragment(t)
	if from == to {
		//seems to be http to https or normalization
		return true
	}
	home := []string{"https://" + fromhost, "https://" + fromhost + "/"}
	for _, loc := range home {
		if loc == to {
			// redirects to home are like 404
			return true
		}
	}
	// error pages
	lowerPath := strings.ToLower(t.Path)
	for _, indicator := range []string{"error", "404", "oops"} {
		if strings.Contains(lowerPath, indicator) {
			return true
		}
	}
	if f1, err := url.QueryUnescape(f.String()); err == nil {
		if f2, err := url.QueryUnescape(t.String()); err == nil {
			if f1 == f2 {
				//ignore % normalization
				return true
			}
		}
	}
	return false
}

func isInterestingRedirect(from, to *url.URL) bool {
	// a query parameter contained the hostname we redirected to
	for _, vals := range from.Query() {
		for _, v := range vals {
			if strings.Contains(v, to.Hostname()) {
				return true
			}
		}
	}
	// to be continued ..?
	return false
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

//https://www.codetd.com/en/article/13918456
func getMaxLengthCommonString(str1, str2 string) string {
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	chs1 := len(runes1)
	chs2 := len(runes2)

	maxLength := 0 //Record the maximum length
	end := 0       //The end position of the record maximum length

	rows := 0
	cols := chs2 - 1
	for rows < chs1 {
		i, j := rows, cols
		length := 0 //record length
		for i < chs1 && j < chs2 {
			if runes1[i] != runes2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}
	return string(runes1[(end - maxLength + 1):(end + 1)])
}
