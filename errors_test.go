package testament_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/blokur/testament"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(count int) string {
	const runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func TestAssertInError(t *testing.T) {
	t.Parallel()
	msg1 := randomString(50)
	msg2 := randomString(60)
	msg3 := randomString(70)
	err1 := errors.New(msg1)
	err2 := errors.New(msg2)
	err3 := errors.New(msg3)
	err1STL := fmt.Errorf("%s %w %s", randomString(80), err1, randomString(90))

	mutiWrap := func(errs ...error) error {
		var err *multierror.Error
		for _, e := range errs {
			err = multierror.Append(err, e)
		}
		return err
	}

	tcs := map[string]struct {
		needle   error
		haystack error
		want     assert.BoolAssertionFunc
	}{
		"same":               {err1, err1, assert.True},
		"different":          {err1, err2, assert.False},
		"nil needle":         {nil, err2, assert.False},
		"nil haystack":       {err1, nil, assert.False},
		"wrapped true":       {err1, errors.Wrap(err1, msg3), assert.True},
		"wrapped false":      {err1, errors.Wrap(err2, msg3), assert.False},
		"wrapped twice":      {err1, errors.Wrap(errors.Wrap(err1, msg3), msg2), assert.True},
		"in wrapped message": {err1, errors.Wrap(err2, err1.Error()), assert.True},
		"stl true":           {err1, err1STL, assert.True},
		"stl false": {err1,
			fmt.Errorf("%s %w %s", randomString(40), err2, randomString(40)),
			assert.False,
		},
		"in stl message": {err1,
			fmt.Errorf("%s %w %s", randomString(40), err2, err1.Error()),
			assert.True},
		"stl(stl)": {err1,
			fmt.Errorf("%s %w %s", randomString(40), err1STL, randomString(40)),
			assert.True,
		},
		"wrap(stl)": {err1,
			errors.Wrap(err1STL, randomString(40)),
			assert.True,
		},
		"stl(wrap)": {err1,
			fmt.Errorf("%s %w %s", randomString(40), errors.Wrap(err1, randomString(40)), randomString(40)),
			assert.True,
		},
		"multi false": {err1, mutiWrap(err2), assert.False},
		"multi true":  {err1, mutiWrap(err1), assert.True},
		"multi mixed": {err1, mutiWrap(err2, err1, err3), assert.True},
		"wrap(multi)": {err1, errors.Wrap(mutiWrap(err2, err1, err3), msg3), assert.True},
		"stl(multi)":  {err1, fmt.Errorf("%s %w %s", msg2, mutiWrap(err2, err1, err3), msg3), assert.True},
		"stl(multi(wrap))": {err1,
			fmt.Errorf("%s %w %s", msg2, mutiWrap(err2, errors.Wrap(err1, randomString(100)), err3), msg3),
			assert.True,
		},
		"stl(wrap(multi))": {err1,
			fmt.Errorf("%s %w %s", msg2, errors.Wrap(mutiWrap(err2, err1, err3), randomString(100)), msg3),
			assert.True,
		},
		"multi(stl(wrap))": {err1,
			mutiWrap(err2, fmt.Errorf("%s %w %s", msg2, errors.Wrap(err1, randomString(100)), err3)),
			assert.True,
		},
		"multi(wrap(stl))": {err1,
			mutiWrap(err2, errors.Wrap(fmt.Errorf("%s %w %s", msg2, err1, err3), randomString(100))),
			assert.True,
		},
		"wrap(multi(stl))": {err1,
			errors.Wrap(mutiWrap(err2, fmt.Errorf("%s %w %s", msg2, err1, err3)), randomString(100)),
			assert.True,
		},
		"wrap(stl(multi))": {err1,
			errors.Wrap(fmt.Errorf("%s %w %s", msg2, mutiWrap(err2, err1, err3), randomString(100)), randomString(100)),
			assert.True,
		},
	}
	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			noT := &testing.T{}
			got := testament.AssertInError(noT, tc.haystack, tc.needle)
			tc.want(t, got)
		})
	}
}
