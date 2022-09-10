package tutils

import "regexp"

func IsAlphanumeric(text string) bool {
	matched, _ := regexp.Match(`/^[a-z0-9]+$/i`, []byte(text))

	return matched
}
