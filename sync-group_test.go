package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testFunction1(ctx context.Context, a int, b *int) error {
	*b = a
	return nil
}

func testFunction2(ctx context.Context, a *int, b *int) error {
	*b = *a
	return nil
}

func TestRun(t *testing.T) {
	group := NewSyncGroup()

	a := 6
	b := 0
	AddMethod(group, testFunction1, OutputArgMethodArgs[int, *int]{
		Context: context.Background(),
		Arg1:    a,
		Arg2:    &b,
	})

	c := 6
	d := 7
	AddMethod(group, testFunction2, OutputArgMethodArgs[*int, *int]{
		Context: context.Background(),
		Arg1:    &c,
		Arg2:    &d,
	})

	require.Equal(t, a, 6)
	require.Equal(t, b, 0)
	require.Equal(t, c, 6)
	require.Equal(t, d, 7)

	err := group.Run()
	require.NoError(t, err)

	require.Equal(t, a, b)
	require.Equal(t, c, d)
}

func TestAddMethod(t *testing.T) {
	group := NewSyncGroup()

	a := 6
	b := 0
	arg := OutputArgMethodArgs[int, *int]{
		Context: context.Background(),
		Arg1:    a,
		Arg2:    &b,
	}
	AddMethod(group, testFunction1, arg)

	require.Equal(t, a, 6)
	require.Equal(t, b, 0)
}
