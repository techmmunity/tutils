package tutils

func NotInSlice[T comparable](baseSlice []T, valuesToCheck []T) []T {
	var inSlice []T

	for _, value := range valuesToCheck {
		if !Contains(baseSlice, value) {
			inSlice = append(inSlice, value)
		}
	}

	return inSlice
}
