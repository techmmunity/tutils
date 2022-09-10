package tutils

func Unique[T comparable](e []T) []T {
	r := []T{}

	for _, s := range e {
			if !Contains(r, s) {
					r = append(r, s)
			}
	}

	return r
}

