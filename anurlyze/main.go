package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var intRegex = regexp.MustCompile(`^[0-9]+$`)
var boolRegex = regexp.MustCompile(`^(true|false|True|False|TRUE|FALSE)$`)
var floatRegex = regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
var sha384Regex = regexp.MustCompile(`^[a-fA-F0-9]{96}$`)
var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)
var sha256Regex = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
var sha224Regex = regexp.MustCompile(`^[a-fA-F0-9]{56}$`)
var sha512Regex = regexp.MustCompile(`^[a-fA-F0-9]{128}$`)
var blake2Regex = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
var whirlpoolRegex = regexp.MustCompile(`^[a-fA-F0-9]{128}$`)
var urlRegex = regexp.MustCompile(`^https?://.+`)

var emptyinterface interface{}

func determineType(input string) string {
	if _, err := strconv.ParseInt(input, 10, 64); err == nil {
		return "Int"
	}

	if boolRegex.MatchString(input) {
		return "Bool"
	}

	if urlRegex.MatchString(input) {
		return "URL"
	}

	if ip := net.ParseIP(input); ip != nil {
		if ip.To4() != nil {
			return "IPv4"
		} else {
			return "IPv6"
		}
	}

	dateFormats := []string{
		time.RFC3339,
		"2006-01-02",
		"02/01/2006",
		"02-01-2006",
		"02.01.2006",
		"02-Jan-2006",
		"02-Jan-06",
		"02-January-2006",
		"02-January-06",
	}

	for _, format := range dateFormats {
		_, err := time.Parse(format, input)
		if err == nil {
			return "Date"
		}
	}

	if _, err := strconv.ParseFloat(input, 64); err == nil && floatRegex.MatchString(input) {
		return "Float"
	}

	if sha1Regex.MatchString(input) {
		return "Hash_SHA-1"
	}

	if sha256Regex.MatchString(input) {
		return "Hash_SHA-256"
	}

	if sha224Regex.MatchString(input) {
		return "Hash_SHA-224"
	}

	if sha384Regex.MatchString(input) {
		return "Hash_SHA-384"
	}

	if sha512Regex.MatchString(input) {
		return "Hash_SHA-512"
	}

	if blake2Regex.MatchString(input) {
		return "Hash_Blake2"
	}

	if whirlpoolRegex.MatchString(input) {
		return "Hash_Whirlpool"
	}

	if md5Regex.MatchString(input) {
		return "Hash_MD5"
	}

	if err := json.Unmarshal([]byte(input), &emptyinterface); err == nil {
		return "JSON"
	}

	if err := xml.Unmarshal([]byte(input), &emptyinterface); err == nil {
		return "XML"
	}

	return ""
}

func main() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read each line from standard input
	for scanner.Scan() {
		line := scanner.Text()

		// Parse the URL from the input line
		u, err := url.Parse(line)
		if err != nil {
			continue
		}

		// Get the query parameters from the URL
		params := u.Query()

		// Create a map to keep track of the types of parameters
		types := map[string]bool{}

		// Check the path segments
		pathSegments := strings.Split(u.Path, "/")
		for _, pathSegment := range pathSegments {
			paramType := determineType(pathSegment)
			if paramType != "" {
				types[paramType] = true
			}
		}

		// Loop through all the parameters in the URL
		for _, values := range params {
			for _, value := range values {
				// Identify the type of each parameter value
				paramType := determineType(value)

				// Add the type to the map if it was successfully identified
				if paramType != "" {
					types[paramType] = true
				}
			}
		}

		// Convert the map to a slice of strings
		typeSlice := []string{}
		for key := range types {
			typeSlice = append(typeSlice, key)
		}

		// Join the slice of types into a comma-separated string
		sort.Strings(typeSlice)
		typesStr := strings.Join(typeSlice, ",")
		if typesStr == "" {
			typesStr = "None"
		}

		// Output the URL and the types of parameters found in it
		fmt.Printf("Url: %s Types: %s\n", line, typesStr)
	}
}
