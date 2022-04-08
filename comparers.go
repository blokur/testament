package testament

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func equal[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	aa := make([]T, len(a))
	copy(aa, a)
	slices.Sort(aa)

	bb := make([]T, len(b))
	copy(bb, b)
	slices.Sort(bb)

	for i := range aa {
		if aa[i] != bb[i] {
			return false
		}
	}
	return true
}

// IntSliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. The slices are copies and can be safely reused.
var IntSliceComparer = cmp.Comparer(func(a, b []int) bool {
	return equal(a, b)
})

// Int32SliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. The slices are copies and can be safely reused.
var Int32SliceComparer = cmp.Comparer(func(a, b []int32) bool {
	return equal(a, b)
})

// Int64SliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. The slices are copies and can be safely reused.
var Int64SliceComparer = cmp.Comparer(func(a, b []int64) bool {
	return equal(a, b)
})

// StringSliceComparer is a go-cmp comparer that doesn't care if the slices in
// question is not sorted. The slices are copies and can be safely reused.
var StringSliceComparer = cmp.Comparer(func(a, b []string) bool {
	return equal(a, b)
})
