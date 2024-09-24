package utils

import "fmt"

func Catch[T any](f func() T) (result T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return f(), nil
}
