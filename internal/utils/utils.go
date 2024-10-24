package utils

import "fmt"

// Catch a panic occuring in the argument function and return it as an error
func Catch[T any](f func() T) (result T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return f(), nil
}
