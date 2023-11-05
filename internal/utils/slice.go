package utils

import "fmt"

func Filter[A any](slice []A, f func(A) bool) []A {
	newValues := []A{}
	for _, value := range slice {
		if f(value) {
			newValues = append(newValues, value)
		}
	}

	return newValues
}

func Map[A any, B any](slice []A, f func(A) B) []B {
	newValues := []B{}
	for _, value := range slice {
		newValues = append(newValues, f(value))
	}

	return newValues
}

func Any[A comparable](slice []A, item A) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func First[A any](slice []A, f func(A) bool) (A, error) {
	for _, value := range slice {
		if f(value) {
			return value, nil
		}
	}
	return *new(A), fmt.Errorf("Item not found")
}
