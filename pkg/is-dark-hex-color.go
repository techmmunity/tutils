package tutils

func IsDarkHexColor(text string) bool {
	return IsHexColor(text) && HexColorLuma(text) < 50;
}
