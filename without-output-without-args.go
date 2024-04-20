package parallelize

import (
	"context"
	"errors"
)

// Use this signature if you want to parallelize a method
//   - Without output
//   - Without arguments
type WithoutOutputWithoutArgsSignature func(context.Context) error

// Output channels for method without output and without args.
type WithoutOutputWithoutArgsChannels struct {
	err chan error
}

// Read output from channels. This method will throw an error if the channels are empty,
// due to the methods not being executed.
func (c WithoutOutputWithoutArgsChannels) Read() error {
	if len(c.err) == 0 {
		return errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.err
}

// Wrapper for method without output and without args.
// The wrapper houses the method, the context and the channels.
type WithoutOutputWithoutArgsWrapper struct {
	method   WithoutOutputWithoutArgsSignature
	ctx      context.Context
	channels WithoutOutputWithoutArgsChannels
}

func newWithoutOutputWithoutArgs(f WithoutOutputWithoutArgsSignature, ctx context.Context) (executable, WithoutOutputWithoutArgsChannels) {
	c := WithoutOutputWithoutArgsChannels{
		make(chan error, 1),
	}
	return &WithoutOutputWithoutArgsWrapper{
		method:   f,
		ctx:      ctx,
		channels: c,
	}, c
}

func (m *WithoutOutputWithoutArgsWrapper) execute() {
	err := m.method(m.ctx)
	m.channels.err <- err
	close(m.channels.err)
}
