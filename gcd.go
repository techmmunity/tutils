package tutils

func GCD(a int, b int) int {
	if b > 0 {
		return GCD(b, a % b)
	}

	return a
}
