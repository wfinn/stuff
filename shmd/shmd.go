package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"flag"
)

func main() {
	syntax := flag.String("s", "sh", "syntax highlighting: string to be put after a beginning `````.")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	linenr := 0
	markdown := false
	waslastamarkdown := true
	insidetruecomment := false
	skip := false
	for scanner.Scan() {
		linenr += 1
		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			fmt.Println(line)
			continue
		}
		if strings.HasPrefix(line, ": '") {
			if line != ": '" {
				fmt.Println(line[3:])
			}
			skip = true
			insidetruecomment = true
		}
		if insidetruecomment && strings.HasPrefix(line, "'") {
			if line != "'" {
				fmt.Println(line[1:])
			}
			skip = true
			insidetruecomment = false
		}
		markdown = strings.HasPrefix(line, "#") || insidetruecomment
		if linenr == 1 && strings.HasPrefix(line, "#!") {
			markdown = false
		}
		if markdown != waslastamarkdown {
			if markdown {
				fmt.Println("```")
			} else {
				fmt.Println("```" + *syntax)
			}
		}
		if markdown && !insidetruecomment {
			line = line[1:]
		}
		if !skip {
			fmt.Println(line)
		}
		waslastamarkdown = markdown
		skip = false
	}

	if !markdown {
		fmt.Println("```")
	}
}
