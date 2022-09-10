package tutils

import "regexp"

func IsIPV4WithMask(text string) bool {
	matched, _ := regexp.Match(`/^(?:[1-9]|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])(?:\.(?:\d|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])){3}\/\d+$/`, []byte(text))

	return matched
}
