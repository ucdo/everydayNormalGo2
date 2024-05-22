package funct

import "golang.org/x/exp/constraints"

// 意味着多态的func

func A[T constraints.Integer](a, b, c T) T {
	return a + b + c
}
