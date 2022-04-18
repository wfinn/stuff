package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

var dedupe bool
var cache map[string]bool

func main() {
	flag.Usage = func() {
		fmt.Printf("%s part\n", os.Args[0])
		fmt.Println("Possible parts: protocol user password hostname port path query fragment ...")
		flag.PrintDefaults()
	}
	flag.BoolVar(&dedupe, "u", false, "don't print dupes")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	if dedupe {
		cache = map[string]bool{}
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if u, err := url.Parse(stdin.Text()); err == nil {
			result := ""
			for _, arg := range flag.Args() {
				result += getPart(u, arg)
			}
			if result != "" {
				println(result)
			}
		}
	}
}

func println(str string) {
	if dedupe {
		if cache[str] {
			return
		} else {
			cache[str] = true
		}
	}
	fmt.Println(str)
}

func getPart(u *url.URL, part string) string {
	switch strings.ToLower(part) {
	case "scheme", "schemes", "proto", "protocol", "protocols":
		return u.Scheme
	case "user", "users", "username", "usernames":
		return u.User.Username()
	case "pass", "password", "passwords":
		pw, _ := u.User.Password()
		return pw
	case "login", "logins", "userinfo", "userpass", "usernamepassword":
		return u.User.String()
	case "domain", "domains", "host", "hosts", "hostname", "hostnames":
		return u.Hostname()
	case "port", "ports":
		return u.Port()
	case "origin", "origins":
		return u.Scheme + "://" + u.Host
	case "path", "paths", "filename", "filenames":
		return u.Path
	case "query", "queries", "params", "parameters":
		return u.RawQuery
	case "fragment", "fragments", "hash", "hashes":
		return u.Fragment
	}
	return part
}
