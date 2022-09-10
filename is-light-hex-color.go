package tutils

func IsLightHexColor(text string) bool {
	return IsHexColor(text) && HexColorLuma(text) >= 50;
}
