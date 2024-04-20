package parallelize

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroup(t *testing.T) {
	var arg interface{} = "test"

	g := NewGroup()
	chans1 := AddWithOutputWithArgs(g, testWithOutputWithArgs, context.Background(), arg)
	chans2 := AddWithOutputWithoutArgs(g, testWithOutputWithoutArgs, context.Background())
	chans3 := AddWithOutputWithArgs(g, testWithOutputWithArgs, context.Background(), arg)
	chans4 := AddWithOutputWithoutArgs(g, testWithOutputWithoutArgs, context.Background())

	g.Execute()

	// comparing outputs with default runs
	out1, err1 := chans1.Read()
	out2, err2 := chans2.Read()
	out3, err3 := chans3.Read()
	out4, err4 := chans4.Read()

	o1, e1 := testWithOutputWithArgs(context.Background(), arg)
	o2, e2 := testWithOutputWithoutArgs(context.Background())
	o3, e3 := testWithOutputWithArgs(context.Background(), arg)
	o4, e4 := testWithOutputWithoutArgs(context.Background())

	require.Equal(t, o1, out1)
	require.Equal(t, o2, out2)
	require.Equal(t, o3, out3)
	require.Equal(t, o4, out4)

	require.Equal(t, e1, err1)
	require.Equal(t, e2, err2)
	require.Equal(t, e3, err3)
	require.Equal(t, e4, err4)
}
