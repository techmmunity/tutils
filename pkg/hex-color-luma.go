package tutils

import (
	"strconv"
	"strings"
)

func getColorWithCorrectLength(color string) string {
	if (len(color) == 3) {
		colorDivided := strings.Split(color, "");

		r := colorDivided[0]
		g := colorDivided[1]
		b := colorDivided[2]

		/**
		 * Duplicate the values, to make the code have length of six
		 */
		colorWithLengthSix := [6]string{r, r, g, g, b, b};

		return strings.Join(colorWithLengthSix[:], "");
	}

	return color;
};

func HexColorLuma(color string) float64 {
	if (!IsHexColor(color)) {
		return 0
	}

	colorWithoutHashtag := strings.Replace(color, "#", "", 1)

	colorWithCorrectLength := getColorWithCorrectLength(colorWithoutHashtag)

	rgb, _ := strconv.ParseInt(colorWithCorrectLength, 16, 64)

	r := float64((rgb >> 16) & 0xff)
	g := float64((rgb >> 8) & 0xff)
	b := float64((rgb >> 0) & 0xff)

	rLuma := 0.2126 * r
	gLuma := 0.7152 * g
	bLuma := 0.0722 * b

	return rLuma + gLuma + bLuma
}
