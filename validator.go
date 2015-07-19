package validator

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	rxURL = regexp.MustCompile(`\b(([\w-]+://?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))`)
)

func IsURL(s string) bool {
	if s == "" || len(s) >= 2083 || len(s) <= 10 || strings.HasPrefix(s, ".") {
		return false
	}
	u, err := url.Parse(s)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	return rxURL.MatchString(s)
}
