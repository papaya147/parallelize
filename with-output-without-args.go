package parallelize

import (
	"context"
	"errors"
)

// Use this signature if you want to parallelize a method
//   - With output
//   - Without arguments
type WithOutputWithoutArgsSignature[O any] func(context.Context) (O, error)

// Output channels for method with output and without args.
type WithOutputWithoutArgsChannels[O any] struct {
	output chan O
	err    chan error
}

// Read output from channels. This method will throw an error if the channels are empty,
// due to the methods not being executed.
func (c WithOutputWithoutArgsChannels[O]) Read() (O, error) {
	if len(c.output) == 0 || len(c.err) == 0 {
		return *new(O), errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.output, <-c.err
}

// Wrapper for method with output and without args.
// The wrapper houses the method, the context and the channels.
type WithOutputWithoutArgsWrapper[O any] struct {
	method   WithOutputWithoutArgsSignature[O]
	ctx      context.Context
	channels WithOutputWithoutArgsChannels[O]
}

func newWithOutputWithoutArgs[O any](f WithOutputWithoutArgsSignature[O], ctx context.Context) (executable, WithOutputWithoutArgsChannels[O]) {
	c := WithOutputWithoutArgsChannels[O]{
		make(chan O, 1),
		make(chan error, 1),
	}
	return &WithOutputWithoutArgsWrapper[O]{
		method:   f,
		ctx:      ctx,
		channels: c,
	}, c
}

func (m *WithOutputWithoutArgsWrapper[O]) execute() {
	out, err := m.method(m.ctx)
	m.channels.output <- out
	m.channels.err <- err
	close(m.channels.output)
	close(m.channels.err)
}
