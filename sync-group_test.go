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
	AddOutputtingMethodWithArgs(group, testFunction1, OutputtingMethodWithArgsParams[int, *int]{
		Context: context.Background(),
		Input:   a,
		Output:  &b,
	})

	c := 6
	d := 7
	AddOutputtingMethodWithArgs(group, testFunction2, OutputtingMethodWithArgsParams[*int, *int]{
		Context: context.Background(),
		Input:   &c,
		Output:  &d,
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
	arg := OutputtingMethodWithArgsParams[int, *int]{
		Context: context.Background(),
		Input:   a,
		Output:  &b,
	}
	AddOutputtingMethodWithArgs(group, testFunction1, arg)

	require.Equal(t, a, 6)
	require.Equal(t, b, 0)
}
