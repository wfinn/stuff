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
var urlRegex = regexp.MustCompile(`^https?://`)

var emptyinterface interface{}

func determineType(input string) string {
	if _, err := strconv.ParseInt(input, 10, 64); err == nil {
		return "Int"
	}

	if boolRegex.MatchString(input) {
		return "Bool"
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

	if _, err := strconv.ParseFloat(input, 64); err == nil {
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

	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		u, err := url.Parse(line)
		if err != nil {
			continue
		}
		params := u.Query()

		types := map[string]bool{}
		for _, values := range params {
			for _, value := range values {
				// Call the determineType function for each parameter value
				paramType := determineType(value)
				if paramType != "" {
					types[paramType] = true
				}
			}
		}
		typeSlice := []string{}

		for key := range types {
			typeSlice = append(typeSlice, key)
		}
		typesStr := strings.Join(typeSlice, ",")
		if typesStr == "" {
			typesStr = "None"
		}

		fmt.Printf("Url: %s Types: %s\n", line, typesStr)
	}
}
