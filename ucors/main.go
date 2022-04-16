package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"

	tld "github.com/jpillora/go-tld"
)

func main() {
	concurrency := flag.Uint("c", 20, "concurrency")
	flag.Parse()
	urls := make(chan string)

	// workers
	var wg sync.WaitGroup
	for i := 0; i < int(*concurrency); i++ {
		wg.Add(1)

		c := getClient()
		go func() {
			defer wg.Done()

			for u := range urls {
				testOrigins(c, u)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		urls <- sc.Text()
	}
	close(urls)

	wg.Wait()

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

func testOrigins(c *http.Client, u string) {

	pp, err := getPermutations(u)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	for _, p := range pp {

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return
		}
		req.Header.Set("Origin", p)

		resp, err := c.Do(req)
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error requesting %s: %s\n", u, err)
			return
		}

		acao := resp.Header.Get("Access-Control-Allow-Origin")
		acac := resp.Header.Get("Access-Control-Allow-Credentials")

		if acao == p {
			fmt.Printf("%s %s %s\n", u, p, acac)
		}
	}
}

func getPermutations(raw string) ([]string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return []string{}, err
	}

	origins := []string{
		"null",
		"https://evil.com",
		"http://evil.com",
	}

	patterns := []string{
		"https://%s.evil.com",
		"https://%sevil.com",
	}
	for i, p := range patterns {
		patterns[i] = fmt.Sprintf(p, u.Hostname())
	}
	origins = append(origins, patterns...)

	if u, err := tld.Parse(raw); err == nil {
		if re, err := regexp.Compile(u.TLD + "$"); err == nil {
			newTLD := "wtf"
			if u.TLD == newTLD {
				newTLD = "ooo"
			}
			origins = append(origins, re.ReplaceAllString(raw, newTLD))
		}
	}
	return origins, nil
}
