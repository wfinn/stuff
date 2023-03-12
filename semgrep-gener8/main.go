package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	lang := flag.String("lang", "generic", "language")
	flag.Parse()
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Printf(`rules:
  - id: change_id_here
    message: change message here
    severity: WARNING
    languages:
      - %s
`, *lang)
	functionNames := make([]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "(")
		functionName := parts[0]
		functionNames[i] = functionName + "(...)"
	}
	fmt.Printf(`    patterns:
    - pattern-either:
      - pattern: %s
`, strings.Join(functionNames, "\n      - pattern: "))

	for _, line := range lines {
		parts := strings.Split(line, "(")
		functionName := parts[0]
		argNumStr := strings.TrimSuffix(parts[1], ")")

		argNum, err := strconv.Atoi(argNumStr)
		if err == nil {
			args := make([]string, argNum+3)
			args[0] = functionName + `(`
			for j := 1; j <= argNum; j++ {
				args[j] = `..., `
			}
			args[argNum+1] = `"...", `
			args[argNum+2] = `...)`
			fmt.Printf("    - pattern-not: %s\n", strings.Join(args, ""))
		}
	}
}
