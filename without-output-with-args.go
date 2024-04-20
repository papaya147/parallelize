package parallelize

import (
	"context"
	"errors"
)

// Use this signature if you want to parallelize a method
//   - Without output
//   - With arguments
type WithoutOutputWithArgsSignature[I any] func(context.Context, I) error

// Output channels for method without output and with args.
type WithoutOutputWithArgsChannels struct {
	err chan error
}

// Read output from channels. This method will throw an error if the channels are empty,
// due to the methods not being executed.
func (c WithoutOutputWithArgsChannels) Read() error {
	if len(c.err) == 0 {
		return errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.err
}

// Wrapper for method without output and with args.
// The wrapper houses the method, the context, the input and the channels.
type WithoutOutputWithArgsWrapper[I any] struct {
	method   WithoutOutputWithArgsSignature[I]
	ctx      context.Context
	input    I
	channels WithoutOutputWithArgsChannels
}

func newWithoutOutputWithArgs[I any](f WithoutOutputWithArgsSignature[I], ctx context.Context, input I) (executable, WithoutOutputWithArgsChannels) {
	c := WithoutOutputWithArgsChannels{
		make(chan error, 1),
	}
	return &WithoutOutputWithArgsWrapper[I]{
		method:   f,
		ctx:      ctx,
		input:    input,
		channels: c,
	}, c
}

func (m *WithoutOutputWithArgsWrapper[I]) execute() {
	m.channels.err <- m.method(m.ctx, m.input)
	close(m.channels.err)
}
