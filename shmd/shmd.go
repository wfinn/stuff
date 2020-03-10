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
	for scanner.Scan() {
		linenr += 1
		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			fmt.Println(line)
			continue
		}
		markdown = strings.HasPrefix(line, "#")
		if linenr == 1 && strings.HasPrefix(line, "#!") {
			markdown = false
		}
		if markdown != waslastamarkdown {
			fmt.Println("```")
		}
		if markdown {
			line = line[1:]
		}
		fmt.Println(line)
		waslastamarkdown = markdown
	}

	if !markdown {
		fmt.Println("```")
	}
}
