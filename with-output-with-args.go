package parallelize

import (
	"context"
	"errors"
)

// Use this signature if you want to parallelize a method
//   - With output
//   - With arguments
type WithOutputWithArgsSignature[I, O any] func(context.Context, I) (O, error)

// Output channels for method with output and with args.
type WithOutputWithArgsChannels[O any] struct {
	output chan O
	err    chan error
}

// Read output from channels. This method will throw an error if the channels are empty,
// due to the methods not being executed.
func (c WithOutputWithArgsChannels[O]) Read() (O, error) {
	if len(c.output) == 0 || len(c.err) == 0 {
		return *new(O), errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.output, <-c.err
}

// Wrapper for method with output and with args.
// The wrapper houses the method, the context, the input and the channels.
type WithOutputWithArgsWrapper[I, O any] struct {
	method   WithOutputWithArgsSignature[I, O]
	ctx      context.Context
	input    I
	channels WithOutputWithArgsChannels[O]
}

func newWithOutputWithArgs[I, O any](f WithOutputWithArgsSignature[I, O], ctx context.Context, input I) (executable, WithOutputWithArgsChannels[O]) {
	c := WithOutputWithArgsChannels[O]{
		make(chan O, 1),
		make(chan error, 1),
	}
	return &WithOutputWithArgsWrapper[I, O]{
		method:   f,
		ctx:      ctx,
		input:    input,
		channels: c,
	}, c
}

func (m *WithOutputWithArgsWrapper[I, O]) execute() {
	out, err := m.method(m.ctx, m.input)
	m.channels.output <- out
	m.channels.err <- err
	close(m.channels.output)
	close(m.channels.err)
}
