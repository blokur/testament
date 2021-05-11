package testament

import (
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// AssertInError returns true if the needle is found in stack, which is created
// either with pkg/errors help, the multierror or Go's error wrap. It will fall
// back to checking the contents of the needle.Error() is in haystack.Error().
func AssertInError(t *testing.T, haystack, needle error) bool {
	t.Helper()
	if haystack == nil || needle == nil {
		t.Errorf("want %v in %v", needle, haystack)
		return false
	}

	if errors.Is(haystack, needle) {
		return true
	}

	contains := func() bool {
		return assert.Containsf(t, haystack.Error(), needle.Error(),
			"want\n\t%v\nin\n\t%v", needle, haystack,
		)
	}

	var merr *multierror.Error
	ok := errors.As(haystack, &merr)
	if !ok {
		return contains()
	}

	for _, err := range merr.Errors {
		if errors.Is(err, needle) {
			return true
		}
	}
	return contains()
}

// RandomString returns a randomly generated string with the length of count.
func RandomString(count int) string {
	const runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, count)
	for i := range b {
		// nolint:gosec // it is supposed to be used in tests.
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
