package tutils

import "regexp"

// Check if has urls in the string
// - https://google.com
// - http://localhost:9090
// - www.google.com
func HasURL(text string) bool {
	matched, _ := regexp.Match(`/(?:(?:https?|ftp|file):\/\/|www\.|ftp\.)[^\s/]*\.\S*/`, []byte(text))

	return matched
}
