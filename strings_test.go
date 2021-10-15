package testament_test

import (
	"strings"
	"testing"
	"testing/quick"

	"github.com/blokur/testament"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()
	t.Run("RandomString", testStringRandomString)
	t.Run("RandomLowerString", testStringRandomLowerString)
	t.Run("StringSlice", testStringStringSlice)
	t.Run("RandomStringSlice", testStringRandomStringSlice)
}

func testStringRandomString(t *testing.T) {
	t.Parallel()
	s1 := testament.RandomString(10)
	assert.Len(t, s1, 10)
	s2 := testament.RandomString(10)
	assert.Len(t, s2, 10)
	assert.NotEqual(t, s1, s2)
	s2 = testament.RandomString(20)
	assert.Len(t, s2, 20)
}

func testStringRandomLowerString(t *testing.T) {
	t.Parallel()
	s1 := testament.RandomLowerString(10)
	assert.Len(t, s1, 10)
	assert.Equal(t, s1, strings.ToLower(s1))
	s2 := testament.RandomLowerString(10)
	assert.Len(t, s2, 10)
	assert.Equal(t, s2, strings.ToLower(s2))
	assert.NotEqual(t, s1, s2)
	s2 = testament.RandomLowerString(20)
	assert.Len(t, s2, 20)
	assert.Equal(t, s2, strings.ToLower(s2))
}

func testStringStringSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.StringSlice(n)
		return len(got) == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func testStringRandomStringSlice(t *testing.T) {
	t.Parallel()
	f := func(m uint8) bool {
		n := int(m)
		if n == 0 {
			n = 10
		}
		got := testament.RandomStringSlice(n)
		return len(got) <= n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
