package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
			skip = true
			insidetruecomment = true
		}
		if insidetruecomment && strings.HasPrefix(line, "'") {
			skip = true
			insidetruecomment = false
		}
		markdown = strings.HasPrefix(line, "#") || insidetruecomment
		if linenr == 1 && strings.HasPrefix(line, "#!") {
			markdown = false
		}
		if markdown != waslastamarkdown {
			fmt.Println("```")
		}
		if markdown {
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
