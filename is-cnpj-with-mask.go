package tutils

import "regexp"

func IsCNPJWithMask(text string) bool {
	matched, _ := regexp.Match(`/^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$/`, []byte(text))

	if (!matched) {
		return false
	}

	notNumbersRegExp, _ := regexp.Compile(`/[^\d]+/g`)

	unmaskedCNPJ := notNumbersRegExp.ReplaceAllString(text, "")

	return IsCNPJ(unmaskedCNPJ)
}
