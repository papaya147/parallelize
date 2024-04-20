package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithOutputWithoutArgs(_ context.Context) (interface{}, error) {
	return "something", nil
}

func TestWithOutputWithoutArgs(t *testing.T) {
	m, c := newWithOutputWithoutArgs(testWithOutputWithoutArgs, context.Background())

	// parallelize run
	m.execute()
	out1, err1 := c.Read()

	// default run
	out2, err2 := testWithOutputWithoutArgs(context.Background())

	require.Equal(t, out1, out2)
	require.Equal(t, err1, err2)
}
