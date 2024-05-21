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
	t.Run("Uint32SliceComparer", testComparerUint32SliceComparer)
	t.Run("Int64SliceComparer", testComparerInt64SliceComparer)
	// t.Run("Uint64SliceComparer", testComparerUint64SliceComparer)
	t.Run("StringSliceComparer", testComparerStringSliceComparer)
}

func testComparerIntSliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandIntSlice(20)
	b := make([]int, len(a))
	copy(b, a)
	b = append(b, rand.Int())

	assert.NotEmpty(t, cmp.Diff(a, b, testament.IntSliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.IntSliceComparer))

	a = []int{3, 1, 2}
	b = []int{3, 1, 2}
	cmp.Diff(a, b, testament.IntSliceComparer)
	assert.Equal(t, []int{3, 1, 2}, b)

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

	a = []int32{3, 1, 2}
	b = []int32{3, 1, 2}
	cmp.Diff(a, b, testament.Int32SliceComparer)
	assert.Equal(t, []int32{3, 1, 2}, b)

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

func testComparerUint32SliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandUint32Slice(20)
	b := make([]uint32, len(a))
	copy(b, a)
	b = append(b, uint32(rand.Int31()))

	assert.NotEmpty(t, cmp.Diff(a, b, testament.Uint32SliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.Uint32SliceComparer))

	a = []uint32{3, 1, 2}
	b = []uint32{3, 1, 2}
	cmp.Diff(a, b, testament.Uint32SliceComparer)
	assert.Equal(t, []uint32{3, 1, 2}, b)

	f := func(a []uint32) bool {
		b := make([]uint32, len(a))
		copy(b, a)
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return cmp.Diff(a, b, testament.Uint32SliceComparer) == ""
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

	a = []int64{3, 1, 2}
	b = []int64{3, 1, 2}
	cmp.Diff(a, b, testament.IntSliceComparer)
	assert.Equal(t, []int64{3, 1, 2}, b)

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

func testComparerStringSliceComparer(t *testing.T) {
	t.Parallel()
	a := testament.RandomStringSlice(20)
	b := make([]string, len(a))
	copy(b, a)
	b = append(b, testament.RandomString(100))

	assert.NotEmpty(t, cmp.Diff(a, b, testament.StringSliceComparer))
	assert.Empty(t, cmp.Diff(a, a, testament.StringSliceComparer))

	a = []string{"d", "a", "b"}
	b = []string{"d", "a", "b"}
	cmp.Diff(a, b, testament.IntSliceComparer)
	assert.Equal(t, []string{"d", "a", "b"}, b)

	f := func(a []string) bool {
		b := make([]string, len(a))
		copy(b, a)
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return cmp.Diff(a, b, testament.StringSliceComparer) == ""
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
