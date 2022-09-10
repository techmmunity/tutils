package tutils

import "regexp"

func HasHtmlTags(text string) bool {
	matched, _ := regexp.Match(`/<[^>]{1,}>/`, []byte(text))

	return matched
}
