package testament

import (
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// GetFreeOpenPort returns a port that is already claimed. It closes the
// listener on test cleanup.
func GetFreeOpenPort(t *testing.T) (uint, net.Listener) {
	t.Helper()
	l, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "could not open a port")
	addr := l.Addr().String()
	parts := strings.Split(addr, ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	require.NoErrorf(t, err, "could not guess the port from %q", addr)
	t.Cleanup(func() {
		l.Close() // nolint:errcheck // not important in tests.
	})
	return uint(port), l
}

// GetFreePort returns a random open port.
func GetFreePort(t *testing.T) uint {
	t.Helper()
	port, l := GetFreeOpenPort(t)
	l.Close() // nolint:errcheck // not important in tests.
	return port
}
