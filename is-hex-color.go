package tutils

import "regexp"

func IsHexColor(text string) bool {
	matched, _ := regexp.Match(`/^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/`, []byte(text))

	return matched
}
