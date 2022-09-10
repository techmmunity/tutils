package tutils

func Contains[T comparable](e []T, c T) bool {
	for _, s := range e {
		if s == c {
			return true
		}
	}
	return false
}
