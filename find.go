package tutils

import (
	"errors"
	"reflect"
)

func Find[T any](arr []T, compare func(T) bool) (T, error) {
	for _, value := range arr {
		if compare(value) {
			return value, nil
		}
	}

	return reflect.Zero(reflect.TypeOf(&arr).Elem().Elem()).Interface().(T), errors.New("not found")
}
