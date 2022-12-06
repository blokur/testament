package testament_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/blokur/testament"
	"github.com/stretchr/testify/require"
)

func TestGetFreeOpenPort(t *testing.T) {
	t.Parallel()
	p, l := testament.GetFreeOpenPort(t)
	require.NotNil(t, l)
	defer l.Close() // nolint:errcheck // not important in testing.

	_, err := net.Listen("tcp", fmt.Sprintf(":%d", p))
	require.Error(t, err)
}

func TestGetFreePort(t *testing.T) {
	t.Parallel()
	p := testament.GetFreePort(t)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", p))
	require.NoError(t, err)
	l.Close() // nolint:errcheck // not important in testing.
}
