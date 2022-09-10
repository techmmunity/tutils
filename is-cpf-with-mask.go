package tutils

import "regexp"

func IsCPFWithMask(text string) bool {
	matched, _ := regexp.Match(`/^\d{3}\.\d{3}\.\d{3}-\d{2}$/`, []byte(text))

	if (!matched) {
		return false
	}

	notNumbersRegExp, _ := regexp.Compile(`/[^\d]+/g`)

	unmaskedCpf := notNumbersRegExp.ReplaceAllString(text, "")

	return IsCPF(unmaskedCpf)
}
