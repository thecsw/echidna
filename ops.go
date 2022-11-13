package gana

import "golang.org/x/exp/constraints"

// Min returns the minimum of two numbers.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two numbers.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Minv returns the minimum of the given values, or if no values are given, the zero value of the type.
func Minv[T constraints.Ordered](vals ...T) T {
	if len(vals) == 0 {
		return ZeroValue[T]()
	}

	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

// Maxv returns the maximum of the given values, or if no values are given, the zero value of the type.
func Maxv[T constraints.Ordered](vals ...T) T {
	if len(vals) == 0 {
		return ZeroValue[T]()
	}

	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// MinMaxv returns the minimum and maximum of the given values, or if no values are given, two zero value of the type.
func MinMaxv[T constraints.Ordered](vals ...T) (T, T) {
	if len(vals) == 0 {
		return ZeroValue[T](), ZeroValue[T]()
	}

	min := vals[0]
	max := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// First returns the first element of the given array, zero value otherwise.
func First[T any](x []T) T {
	if len(x) == 0 {
		return ZeroValue[T]()
	}
	return x[0]
}

// Last returns the last element of the given array, zero value otherwise.
func Last[T any](x []T) T {
	l := len(x)
	if l == 0 {
		return ZeroValue[T]()
	}
	return x[l-1]
}

// ZeroValue returns the zero value of any type.
func ZeroValue[T any]() T {
	var t T
	return t
}

// Map calls the function on every array element and returns results in list.
func Map[T, S any](f func(T) S, arr []T) []S {
	what := make([]S, len(arr))
	for i, v := range arr {
		what[i] = f(v)
	}
	return what
}

// Filter calls the function and returns list of values that returned true.
func Filter[T any](f func(T) bool, arr []T) []T {
	what := make([]T, 0, len(arr))
	for _, v := range arr {
		if !f(v) {
			continue
		}
		what = append(what, v)
	}
	return what
}

// Take returns up to the first `num` elements.
func Take[T any](num int, arr []T) []T {
	if len(arr) < num {
		return arr
	}
	return arr[:num]
}

// Tail returns up to the last `num` elements.
func Tail[T any](num int, arr []T) []T {
	if len(arr) < num {
		return arr
	}
	return arr[len(arr)-num:]
}

// Drop allocates a new slice, with the first `num` elements dropped.
func Drop[T any, U constraints.Unsigned](num U, arr []T) []T {
	l := U(len(arr))
	if l < num {
		return []T{}
	}

	slice := make([]T, l-num)
	copy(slice, arr[num:])

	return slice
}

// Skips skips the first `num` elements by slicing (underlying array unaffected).
func Skip[T any, U constraints.Unsigned](num U, arr []T) []T {
	if U(len(arr)) < num {
		return arr[:0]
	}
	return arr[num:]
}

// Any returns true if any element in the list matches the given value.
func Any[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

// Anyf returns true if any elemens in the list returns true when passed to foo.
func Anyf[T any](foo func(v T) bool, arr []T) bool {
	for _, v := range arr {
		if foo(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the list match the given value.
func All[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if val != v {
			return false
		}
	}
	return true
}

// Allf returns true if all elements in the array return true when passed to foo.
func Allf[T any](foo func(v T) bool, arr []T) bool {
	for _, v := range arr {
		if !foo(v) {
			return false
		}
	}
	return true
}

// Repeat creates a list of given size consisting of the same given value.
func Repeat[T any](val T, size int) []T {
	arr := make([]T, size)
	for i := range arr {
		arr[i] = val
	}
	return arr
}

// GetPointer returns the pointer of a given lhs.
func GetPointer[T any](val T) *T {
	return &val
}
