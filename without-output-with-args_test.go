package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithoutOutputWithArgs(_ context.Context, arg interface{}) error {
	return nil
}

func TestWithoutOutputWithArgs(t *testing.T) {
	var arg interface{} = "test"
	m, c := newWithoutOutputWithArgs(testWithoutOutputWithArgs, context.Background(), arg)

	// parallelize run
	m.execute()
	err1 := c.Read()

	// default run
	err2 := testWithoutOutputWithArgs(context.Background(), arg)

	require.Equal(t, err1, err2)
}
