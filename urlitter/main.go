package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var maxRunes = flag.Int("max-runes", 1000, "The maximum number of runes allowed in a URL")
var expectedCharsRatio = flag.Float64("expected-ratio", 0.5, "The ratio of expected characters in a URL")

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		u := scanner.Text()

		// Filter URLs that have an invalid format or a non-supported protocol
		parsedURL, err := url.Parse(u)
		if err != nil || parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			continue
		}

		// Filter URLs whose path contains "http://" or "https://"
		if strings.Contains(parsedURL.Path, "http://") || strings.Contains(parsedURL.Path, "https://") {
			continue
		}

		// Compile the regular expression for matching expected characters
		expectedCharsRegexp := regexp.MustCompile(`[a-fA-F0-9.-/]*`)

		// Count the number of expected characters in the URL path
		expectedChars := expectedCharsRegexp.FindAllString(parsedURL.Path, -1)
		numExpectedChars := 0
		for _, c := range expectedChars {
			numExpectedChars += len(c)
		}

		// Filter URLs that are longer than the maximum number of runes allowed
		// or have a ratio of expected characters to total characters in the URL path
		// less than the specified ratio
		if utf8.RuneCountInString(u) > *maxRunes ||
			len(parsedURL.Path) == 0 ||
			float64(numExpectedChars)/float64(len(parsedURL.Path)) < *expectedCharsRatio {
			continue
		}

		// Only print the URL if it has passed all the filters
		fmt.Println(u)
	}
}
