package tutils

import "strconv"

func IsCNPJ(text string) bool {
	notCnpj := []string{
		"00000000000000",
		"11111111111111",
		"22222222222222",
		"33333333333333",
		"44444444444444",
		"55555555555555",
		"66666666666666",
		"77777777777777",
		"88888888888888",
		"99999999999999",
	}

	if (text == "") {
		return false
	}

	if (len(text) != 14) {
		return false
	}

	if (Contains(notCnpj, text)) {
		return false
	}

	length := len(text) - 2
	numbers := text[:length]
	digits := text[length:]
	sum := int64(0)
	pos := length - 7

	for i := length; i >= 1; i-- {
		nbr, _ := strconv.ParseInt(string(numbers[length - i]), 10, 64)

		pos--

		sum += nbr * int64(pos)

		if (pos < 2) {
			pos = 9
		}
	}

	var result int64

	if (sum % 11 < 2) {
		result = 0
	} else {
		result = 11 - (sum % 11)
	}

	firstChar, _ := strconv.ParseInt(string(digits[0]), 10, 64)

	if result != firstChar {
		return false
	}

	length += 1
	numbers = text[0:length]
	sum = 0
	pos = length - 7

	for i := length; i >= 1; i-- {
		nbr, _ := strconv.ParseInt(string(numbers[length-i]), 10, 64)

		pos--

		sum += nbr * int64(pos)

		if (pos < 2) {
			pos = 9
		}
	}

	if (sum % 11 < 2) {
		result = 0
	} else {
		result = 11 - (sum % 11)
	}

	secondChar, _ := strconv.ParseInt(string(digits[1]), 10, 64)

	return result == secondChar;
}
