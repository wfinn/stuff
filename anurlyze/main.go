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
var langregex = regexp.MustCompile("(?i)^(af|af[-_]ZA|ar|ar[-_]AE|ar[-_]BH|ar[-_]DZ|ar[-_]EG|ar[-_]IQ|ar[-_]JO|ar[-_]KW|ar[-_]LB|ar[-_]LY|ar[-_]MA|ar[-_]OM|ar[-_]QA|ar[-_]SA|ar[-_]SY|ar[-_]TN|ar[-_]YE|az|az[-_]AZ|az[-_]AZ|be|be[-_]BY|bg|bg[-_]BG|bs[-_]BA|ca|ca[-_]ES|cs|cs[-_]CZ|cy|cy[-_]GB|da|da[-_]DK|de|de[-_]AT|de[-_]CH|de[-_]DE|de[-_]LI|de[-_]LU|dv|dv[-_]MV|el|el[-_]GR|en|en[-_]AU|en[-_]BZ|en[-_]CA|en[-_]CB|en[-_]GB|en[-_]IE|en[-_]JM|en[-_]NZ|en[-_]PH|en[-_]TT|en[-_]US|en[-_]ZA|en[-_]ZW|eo|es|es[-_]AR|es[-_]BO|es[-_]CL|es[-_]CO|es[-_]CR|es[-_]DO|es[-_]EC|es[-_]ES|es[-_]ES|es[-_]GT|es[-_]HN|es[-_]MX|es[-_]NI|es[-_]PA|es[-_]PE|es[-_]PR|es[-_]PY|es[-_]SV|es[-_]UY|es[-_]VE|et|et[-_]EE|eu|eu[-_]ES|fa|fa[-_]IR|fi|fi[-_]FI|fo|fo[-_]FO|fr|fr[-_]BE|fr[-_]CA|fr[-_]CH|fr[-_]FR|fr[-_]LU|fr[-_]MC|gl|gl[-_]ES|gu|gu[-_]IN|he|he[-_]IL|hi|hi[-_]IN|hr|hr[-_]BA|hr[-_]HR|hu|hu[-_]HU|hy|hy[-_]AM|id|id[-_]ID|is|is[-_]IS|it|it[-_]CH|it[-_]IT|ja|ja[-_]JP|ka|ka[-_]GE|kk|kk[-_]KZ|kn|kn[-_]IN|ko|ko[-_]KR|kok|kok[-_]IN|ky|ky[-_]KG|lt|lt[-_]LT|lv|lv[-_]LV|mi|mi[-_]NZ|mk|mk[-_]MK|mn|mn[-_]MN|mr|mr[-_]IN|ms|ms[-_]BN|ms[-_]MY|mt|mt[-_]MT|nb|nb[-_]NO|nl|nl[-_]BE|nl[-_]NL|nn[-_]NO|ns|ns[-_]ZA|pa|pa[-_]IN|pl|pl[-_]PL|ps|ps[-_]AR|pt|pt[-_]BR|pt[-_]PT|qu|qu[-_]BO|qu[-_]EC|qu[-_]PE|ro|ro[-_]RO|ru|ru[-_]RU|sa|sa[-_]IN|se|se[-_]FI|se[-_]FI|se[-_]FI|se[-_]NO|se[-_]NO|se[-_]NO|se[-_]SE|se[-_]SE|se[-_]SE|sk|sk[-_]SK|sl|sl[-_]SI|sq|sq[-_]AL|sr[-_]BA|sr[-_]BA|sr[-_]SP|sr[-_]SP|sv|sv[-_]FI|sv[-_]SE|sw|sw[-_]KE|syr|syr[-_]SY|ta|ta[-_]IN|te|te[-_]IN|th|th[-_]TH|tl|tl[-_]PH|tn|tn[-_]ZA|tr|tr[-_]TR|tt|tt[-_]RU|ts|uk|uk[-_]UA|ur|ur[-_]PK|uz|uz[-_]UZ|uz[-_]UZ|vi|vi[-_]VN|xh|xh[-_]ZA|zh|zh[-_]CN|zh[-_]HK|zh[-_]MO|zh[-_]SG|zh[-_]TW|zu|zu[-_]zA)$")

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

	if langregex.MatchString(input) {
		return "Langcode"
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
