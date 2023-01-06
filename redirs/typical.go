package main

import (
	"net/url"
	"strings"
)

func looksLikeRedirect(urlStr string) bool {
	if u, err := url.Parse(urlStr); err == nil {
		for k, vals := range u.Query() {
			if isKnownRedirParam(k) {
				return true
			}
			for _, v := range vals {
				if looksLikeURI(v) {
					return true
				}
			}
		}
		if looksLikeRedirectPath(u.Path) {
			return true
		}
	}
	return false
}

func isKnownRedirParam(paramName string) bool {
	for _, name := range paramNames {
		if name == paramName {
			return true
		}
	}
	return false
}

func looksLikeURI(paramVal string) bool {
	return strings.HasPrefix("/", paramVal) || strings.HasPrefix("http:", paramVal) || strings.HasPrefix("https:", paramVal)
}

// looksLikeRedirectPath currently doesn't have it's own list of identifiers and works with "paramNames"
func looksLikeRedirectPath(path string) bool {
	for _, segment := range strings.Split(path, "/") {
		for _, redirName := range paramNames {
			if segment == redirName {
				return true
			}
		}
	}
	return false
}
