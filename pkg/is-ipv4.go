package tutils

import "regexp"

func IsIPV4(text string) bool {
	matched, _ := regexp.Match(`/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/`, []byte(text))

	return matched
}
