package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

var seen = map[string]bool{}
var printdupes = false

func main() {
	printkeys := flag.Bool("keys", false, "print parameter keys")
	printvals := flag.Bool("vals", false, "print parameter vals")
	flag.BoolVar(&printdupes, "dupes", false, "disable deduplication")
	flag.Parse()
	if *printkeys == false && *printvals == false {
		flag.PrintDefaults()
		return
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if u, err := url.Parse(stdin.Text()); err == nil {
			for k, vals := range u.Query() {
				if *printkeys {
					print(k)
				}
				if *printvals {
					for _, v := range vals {
						print(v)
					}
				}
			}
		}
	}
}

func print(text string) {
	if !printdupes {
		if seen[text] {
			return
		} else {
			seen[text] = true
		}
	}
	fmt.Println(text)
}
