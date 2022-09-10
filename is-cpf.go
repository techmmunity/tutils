package tutils

import (
	"fmt"
	"strconv"
)

func IsCPF(text string) bool {
	notCpf := []string{
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

	if (Contains(notCpf, text)) {
		return false
	}

	temp := text[0:9]
	count := int64(10)
	total := int64(0)

	for _, nbr := range temp {
		nbrInt, _ := strconv.ParseInt(string(nbr), 10, 64)

		total += nbrInt * count

		count--
	}

	total = (total * 10) % 11

	if (total > 9) {
		total = 0
	}

	if fmt.Sprint(total) != string(text[9]) {
		return false
	}

	temp = text[0:10]
	count = 11
	total = 0

	for _, nbr := range temp {
		nbrInt, _ := strconv.ParseInt(string(nbr), 10, 64)

		total += nbrInt * count

		count--
	}

	total = (total * 10) % 11

	if (total > 9) {
		total = 0
	}

	return fmt.Sprint(total) == string(text[10])
}
