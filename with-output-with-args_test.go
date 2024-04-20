package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithOutputWithArgs(_ context.Context, arg interface{}) (interface{}, error) {
	return arg, nil
}

func TestWithOutputWithArgs(t *testing.T) {
	var arg interface{} = "test"
	m, c := newWithOutputWithArgs(testWithOutputWithArgs, context.Background(), arg)

	// parallelize run
	m.execute()
	out1, err1 := c.Read()

	// default run
	out2, err2 := testWithOutputWithArgs(context.Background(), arg)

	require.Equal(t, out1, out2)
	require.Equal(t, err1, err2)
}
