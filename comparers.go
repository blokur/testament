package testament

import (
	"sort"

	"github.com/google/go-cmp/cmp"
)

// IntSliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. Note that this will sort the slices in place.
var IntSliceComparer = cmp.Comparer(func(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })
	sort.SliceStable(b, func(i, j int) bool { return b[i] < b[j] })
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
})

// Int32SliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. Note that this will sort the slices in place.
var Int32SliceComparer = cmp.Comparer(func(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })
	sort.SliceStable(b, func(i, j int) bool { return b[i] < b[j] })
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
})

// Int64SliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. Note that this will sort the slices in place.
var Int64SliceComparer = cmp.Comparer(func(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })
	sort.SliceStable(b, func(i, j int) bool { return b[i] < b[j] })
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
})
