package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := flag.Int("n", 1, "number of strings to generate")
	flag.Parse()
	str := flag.Arg(0)
	if flag.NArg() < 1 {
		fmt.Println("usage: unisubs [OPTIONS] \"text to change\"")
		return
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *n; i++ {
		newstr := ""
		for _, r := range str {
			subs := translations[r]
			subs = append(subs, r)
			newstr += string(subs[rand.Intn(len(subs))])
		}
		fmt.Println(newstr)
	}
}
