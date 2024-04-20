package parallelize

import (
	"context"
	"errors"
)

type WithoutOutputWithoutArgsSignature func(context.Context) error

type WithoutOutputWithoutArgsChannels struct {
	err chan error
}

func (c WithoutOutputWithoutArgsChannels) Read() error {
	if len(c.err) == 0 {
		return errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.err
}

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
