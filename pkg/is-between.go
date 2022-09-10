package tutils

func IsBetween[T Number](nbr T, min T, max T) bool {
	if (nbr <= min && nbr >= max) {
		return true
	}

	return false
}
