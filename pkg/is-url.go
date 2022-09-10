package tutils

import "regexp"

func IsURL(text string) bool {
	matched, _ := regexp.Match(`/^(?:(?:https?|ftp|file):\/\/|www\.|ftp\.)[^\s/]*\.\S*$/`, []byte(text))

	return matched
}
