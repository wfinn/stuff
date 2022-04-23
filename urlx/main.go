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
		fmt.Println("Example: urlx --get -u hosts < urls.txt")
		fmt.Println("Possible parts: protocol user password hostname port path query fragment ...")
		flag.PrintDefaults()
	}
	flag.BoolVar(&dedupe, "u", false, "don't print dupes")
	get := flag.Bool("get", false, "extract parts from urls")
	set := flag.Bool("set", false, "modify parts of urls")
	flag.Parse()
	if flag.NArg() < 1 || (*set == false && *get == false) {
		flag.Usage()
		return
	}
	if dedupe {
		cache = map[string]bool{}
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if u, err := url.Parse(stdin.Text()); err == nil {
			if *set {
				for _, arg := range flag.Args() {
					key, value := arg, ""
					if kv := strings.SplitN(arg, "=", 2); kv != nil {
						key = mapPart(kv[0])
						value = kv[1]
					}
					setPart(u, key, value)
				}
				println(u.String())
			} else if *get {
				result := ""
				for _, arg := range flag.Args() {
					result += getPart(u, mapPart(arg))
				}
				if result != "" {
					println(result)
				}
			}
		}
	}
}

func getPart(u *url.URL, part string) string {
	switch part {
	case "scheme":
		return u.Scheme
	case "user":
		return u.User.Username()
	case "pass":
		pw, _ := u.User.Password()
		return pw
	case "login":
		return u.User.String()
	case "domain":
		return u.Hostname()
	case "port":
		return u.Port()
	case "origin":
		return u.Scheme + "://" + u.Host
	case "path":
		return u.Path
	case "query":
		return u.RawQuery
	case "hash":
		return u.Fragment
	}
	return part
}

func setPart(u *url.URL, part, value string) *url.URL {
	value = strings.ReplaceAll(value, "{{value}}", getPart(u, part))
	switch part {
	case "scheme":
		u.Scheme = value
	case "user":
		if pw, haspw := u.User.Password(); haspw {
			u.User = url.UserPassword(value, pw)
		} else {
			u.User = url.User(value)
		}
	case "pass":
		u.User = url.UserPassword(u.User.Username(), value)
	case "login":
		userpass := strings.SplitN(value, ":", 2)
		if len(userpass) > 1 {
			u.User = url.UserPassword(userpass[0], userpass[1])
		} else {
			u.User = url.User(value)
		}
	case "domain":
		u.Host = value
	case "port":
		u.Host = u.Hostname() + ":" + value
	case "origin":
		split := strings.SplitN(value, "://", 2)
		if len(split) > 1 {
			u.Scheme = split[0]
			u.Host = split[1]
		}
	case "path":
		u.Path = value
	case "query":
		u.RawQuery = value
	case "hash":
		u.Fragment = value
	}
	return u
}

// This should make it easy to use this tool, don't have to remember any terms
func mapPart(part string) string {
	switch strings.ToLower(part) {
	case "scheme", "schemes", "proto", "protocol", "protocols":
		return "scheme"
	case "user", "users", "username", "usernames":
		return "user"
	case "pass", "password", "passwords":
		return "pass"
	case "login", "logins", "userinfo", "userpass", "usernamepassword":
		return "login"
	case "domain", "domains", "host", "hosts", "hostname", "hostnames":
		return "domain"
	case "port", "ports":
		return "port"
	case "origin", "origins":
		return "origin"
	case "path", "paths", "filename", "filenames":
		return "path"
	case "query", "queries", "params", "parameters", "search":
		return "query"
	case "fragment", "fragments", "hash", "hashes":
		return "hash"
	}
	return part
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
