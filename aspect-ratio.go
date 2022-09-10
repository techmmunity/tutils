package tutils

import "fmt"

func AspectRatio(width int, height int) string {
	divisor := GCD(width, height)

	return fmt.Sprintf("%d:%d", width / divisor, height / divisor)
}
