package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithoutOutputWithoutArgs(_ context.Context) error {
	return nil
}

func TestWithoutOutputWithoutArgs(t *testing.T) {
	m, c := newWithoutOutputWithoutArgs(testWithoutOutputWithoutArgs, context.Background())

	// parallelize run
	m.execute()
	err1 := c.Read()

	// default run
	err2 := testWithoutOutputWithoutArgs(context.Background())

	require.Equal(t, err1, err2)
}
