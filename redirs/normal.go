package main

import (
	"net/url"
	"strings"
)

func normalHost(u *url.URL) string {
	return strings.ToLower(wwwregex.ReplaceAllString(u.Hostname(), ""))
}

func normalPath(u *url.URL) string {
	path := u.EscapedPath()
	if path == "" {
		return "/"
	}
	return path
}

func normalQuery(u *url.URL) string {
	query := u.Query().Encode()
	if query == "" {
		return ""
	}
	return "?" + query
}

func normalFragment(u *url.URL) string {
	if u.Fragment == "" {
		return ""
	}
	return "#" + u.Fragment
}
