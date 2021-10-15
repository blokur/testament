package testament_test

import (
	"math/rand"
	"testing"
	"testing/quick"

	"github.com/blokur/testament"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestComparer(t *testing.T) {
	t.Parallel()
	t.Run("IntSliceComparer", testComparerIntSliceComparer)
	t.Run("Int32SliceComparer", testComparerInt32SliceComparer)
	t.Run("Int64SliceComparer", testComparerInt64SliceComparer)
}

func testComparerIntSliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandIntSlice(20)
	b := make([]int, len(a))
	copy(b, a)
	b = append(b, rand.Int())

	assert.NotEmpty(t, cmp.Diff(a, b, testament.IntSliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.IntSliceComparer))

	f := func(a []int) bool {
		b := make([]int, len(a))
		copy(b, a)
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return cmp.Diff(a, b, testament.IntSliceComparer) == ""
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testComparerInt32SliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandInt32Slice(20)
	b := make([]int32, len(a))
	copy(b, a)
	b = append(b, rand.Int31())

	assert.NotEmpty(t, cmp.Diff(a, b, testament.Int32SliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.Int32SliceComparer))

	f := func(a []int32) bool {
		b := make([]int32, len(a))
		copy(b, a)
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return cmp.Diff(a, b, testament.Int32SliceComparer) == ""
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testComparerInt64SliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandInt64Slice(20)
	b := make([]int64, len(a))
	copy(b, a)
	b = append(b, rand.Int63())

	assert.NotEmpty(t, cmp.Diff(a, b, testament.Int64SliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.Int64SliceComparer))

	f := func(a []int64) bool {
		b := make([]int64, len(a))
		copy(b, a)
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return cmp.Diff(a, b, testament.Int64SliceComparer) == ""
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
