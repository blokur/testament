package testament_test

import (
	"testing"
	"testing/quick"

	"github.com/blokur/testament"
)

func TestSlices(t *testing.T) {
	t.Parallel()
	t.Run("IntSlice", testSlicesIntSlice)
	t.Run("UintSlice", testSlicesUintSlice)
	t.Run("RandIntSlice", testSlicesRandIntSlice)
	t.Run("RandUintSlice", testSlicesRandUintSlice)
	t.Run("Int32Slice", testSlicesInt32Slice)
	t.Run("Uint32Slice", testSlicesUint32Slice)
	t.Run("RandInt32Slice", testSlicesRandInt32Slice)
	t.Run("RandUint32Slice", testSlicesRandUint32Slice)
	t.Run("Int64Slice", testSlicesInt64Slice)
	t.Run("Uint64Slice", testSlicesUint64Slice)
	t.Run("RandInt64Slice", testSlicesRandInt64Slice)
	t.Run("RandUint64Slice", testSlicesRandUint64Slice)
}

func testSlicesIntSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.IntSlice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesUintSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.UintSlice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandIntSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandIntSlice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandUintSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandUintSlice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesInt32Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.Int32Slice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesUint32Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.Uint32Slice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandInt32Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandInt32Slice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandUint32Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandUint32Slice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesInt64Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.Int64Slice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesUint64Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.Uint64Slice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandInt64Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandInt64Slice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testSlicesRandUint64Slice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandUint64Slice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
