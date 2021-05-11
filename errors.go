package testament

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// AssertIsCode is a helper to assert the err error contains the code.
func AssertIsCode(t *testing.T, err error, code codes.Code) bool {
	t.Helper()
	if !assert.Error(t, err, "empty error") {
		return false
	}
	st := status.Convert(err)
	if !assert.NotNil(t, st, "empty status") {
		return false
	}
	return assert.EqualValuesf(t, code, st.Code(),
		"got %q code, want %q. error: %#v", st.Code(), code, err,
	)
}

// GetFreeOpenPort returns a port that is already claimed.
func GetFreeOpenPort(t *testing.T) (uint, net.Listener) {
	t.Helper()
	l, err := net.Listen("tcp", ":0") //nolint:gosec // this is used in tests.
	require.NoError(t, err, "could not open a port")
	addr := l.Addr().String()
	parts := strings.Split(addr, ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	require.NoErrorf(t, err, "could not guess the port from %q", addr)
	return uint(port), l
}

// GetFreePort returns a random open port.
func GetFreePort(t *testing.T) uint {
	t.Helper()
	port, l := GetFreeOpenPort(t)
	l.Close() // nolint // not impoerant in tests.
	return port
}
