package testament

import (
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// AssertInError returns true if the needle is found in stack, which is
// created either with pkg/errors help, the multierror or Go's error wrap. It
// will fall back to checking the contents of the needle.Error() is in
// haystack.Error().
func AssertInError(t *testing.T, haystack, needle error) bool {
	t.Helper()
	if haystack == nil || needle == nil {
		t.Errorf("want %v in %v", needle, haystack)
		return false
	}
	if errors.Cause(haystack) == needle {
		return true
	}
	if errors.Is(haystack, needle) {
		return true
	}
	contains := func() bool {
		return assert.Containsf(t, haystack.Error(), needle.Error(),
			"want\n\t%v\nin\n\t%v", needle, haystack,
		)
	}
	merr, ok := errors.Cause(haystack).(*multierror.Error)
	if !ok {
		return contains()
	}
	for _, err := range merr.Errors {
		if errors.Cause(err) == needle {
			return true
		}
	}
	return contains()
}
