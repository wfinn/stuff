package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	// Accept single input as argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <unicode_point>")
		return
	}

	// Parse the input as a unicode point
	unicodePoint := os.Args[1]
	regex := regexp.MustCompile(`(?i)(u\+|\\u|%u|\+)?([0-9a-f]{4,6})`)
	match := regex.FindStringSubmatch(unicodePoint)

	if len(match) != 3 {
		fmt.Println("Error: Invalid unicode point format")
		return
	}

	code, err := strconv.ParseInt(match[2], 16, 32)
	if err != nil {
		fmt.Println("Error: Invalid unicode point format")
		return
	}

	// Convert the unicode point to a string and copy it to the clipboard
	unicodeChar := string(code)
	err = clipboard.WriteAll(unicodeChar)
	if err != nil {
		fmt.Println("Error: Failed to copy to clipboard")
		return
	}

	fmt.Println("Copied to clipboard:", unicodeChar)
}
