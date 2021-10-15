package testament

import "math/rand"

// IntSlice returns a slice of int elements with the provided length.
func IntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

// RandIntSlice returns a slice of int elements with random length of less than
// n.
func RandIntSlice(n int) []int {
	s := make([]int, rand.Intn(n))
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

// Int32Slice returns a slice of int32 elements with provided length.
func Int32Slice(n int) []int32 {
	s := make([]int32, n)
	for i := range s {
		s[i] = rand.Int31()
	}
	return s
}

// RandInt32Slice returns a slice of int32 elements with random length of less
// than n.
func RandInt32Slice(n int) []int32 {
	s := make([]int32, rand.Intn(n))
	for i := range s {
		s[i] = rand.Int31()
	}
	return s
}

// Int64Slice returns a slice of int64 elements with provided length.
func Int64Slice(n int) []int64 {
	s := make([]int64, n)
	for i := range s {
		s[i] = rand.Int63()
	}
	return s
}

// RandInt64Slice returns a slice of int64 elements with random length of less
// than n.
func RandInt64Slice(n int) []int64 {
	s := make([]int64, rand.Intn(n))
	for i := range s {
		s[i] = rand.Int63()
	}
	return s
}
