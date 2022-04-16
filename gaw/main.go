package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

var cache map[string]bool
var min uint
var max uint
var splitregex = regexp.MustCompile("(<|>|\"|'|\\.|-|,|;|:|/|\\(|\\)|{|}|_)")

func main() {
	flag.UintVar(&min, "min", 3, "minimum length")
	flag.UintVar(&max, "max", 20, "maximum length (0 = no maximum)")
	flag.Usage = func() {
		fmt.Printf("Usage: %s < urls.txt > wordlist.txt\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	cache = map[string]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if u, err := url.Parse(scanner.Text()); err == nil {
			printWord(u.User.Username())
			if pw, haspw := u.User.Password(); haspw {
				printWord(pw)
			}
			printWord(u.Hostname())
			printWords(strings.Split(u.Path, "/"))
			printWord(u.Fragment)
			for k, vals := range u.Query() {
				printWord(k)
				printWords(vals)
			}
		} else {
			log.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printWord(str string) {
	_printUnique(str)
	_printUnique(strings.TrimSuffix(str, path.Ext(str)))
	for _, s := range splitregex.Split(str, -1) {
		_printUnique(s)
	}
}

func printWords(strs []string) {
	for _, str := range strs {
		printWord(str)
	}
}

func _printUnique(str string) {
	l := uint(len(str))
	if l >= min && (max == 0 || l <= max) && strings.TrimSpace(str) != "" && !cache[str] {
		cache[str] = true
		fmt.Println(str)
	}
}
