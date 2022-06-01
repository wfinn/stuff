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
var seen = map[string]bool{}

func main() {
	flag.Usage = func() {
		fmt.Printf("%s part\n", os.Args[0])
		fmt.Println("Possible parts: protocol user password hostname port path query fragment ...")
		fmt.Println("")
		fmt.Println("--keys or --vals are special cases as they will print more than one line per url")
		flag.PrintDefaults()
	}
	flag.BoolVar(&dedupe, "u", false, "don't print dupes")
	printkeys := flag.Bool("keys", false, "print list of all query keys")
	printvals := flag.Bool("vals", false, "print list of all query vals")
	flag.Parse()
	if flag.NArg() < 1 && !*printkeys && !*printvals {
		flag.Usage()
		return
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if u, err := url.Parse(stdin.Text()); err == nil {
			if *printkeys || *printvals {
				for k, vals := range u.Query() {
					if *printkeys {
						println(k)
					}
					if *printvals {
						for _, v := range vals {
							println(v)
						}
					}
				}
			}
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
	case "query", "queries", "params", "parameters", "search", "searches":
		return u.RawQuery
	case "fragment", "fragments", "hash", "hashes":
		return u.Fragment
	}
	return part
}

func println(str string) {
	if dedupe {
		if seen[str] {
			return
		} else {
			seen[str] = true
		}
	}
	fmt.Println(str)
}
