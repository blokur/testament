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

// UintSlice returns a slice of uint elements with the provided length.
func UintSlice(n int) []uint {
	s := make([]uint, n)
	for i := range s {
		// We know the returned value of rand.Int() is positive.
		s[i] = uint(rand.Int())
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

// RandUintSlice returns a slice of uint elements with random length of less
// than n.
func RandUintSlice(n int) []uint {
	s := make([]uint, rand.Intn(n))
	for i := range s {
		// We know the returned value of rand.Int() is positive.
		s[i] = uint(rand.Int())
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

// Uint32Slice returns a slice of uint32 elements with provided length.
func Uint32Slice(n int) []uint32 {
	s := make([]uint32, n)
	for i := range s {
		// We know the returned value of rand.Int31() is positive.
		s[i] = uint32(rand.Int31())
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

// RandUint32Slice returns a slice of uint32 elements with random length of
// less than n.
func RandUint32Slice(n int) []uint32 {
	// We already know the return value of rand.Intn() is a positive value.
	s := make([]uint32, rand.Intn(n))
	for i := range s {
		s[i] = uint32(rand.Int31())
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

// Uint64Slice returns a slice of uint64 elements with provided length.
func Uint64Slice(n int) []uint64 {
	s := make([]uint64, n)
	for i := range s {
		// We know the returned value of rand.Int63() is positive.
		s[i] = uint64(rand.Int63())
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

// RandUint64Slice returns a slice of uint64 elements with random length of
// less than n.
func RandUint64Slice(n int) []uint64 {
	s := make([]uint64, rand.Intn(n))
	for i := range s {
		// We know the returned value of rand.Int63() is positive.
		s[i] = uint64(rand.Int63())
	}
	return s
}
