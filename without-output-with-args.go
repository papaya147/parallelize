package parallelize

import (
	"context"
	"errors"
)

type WithoutOutputWithArgsSignature[I any] func(context.Context, I) error

type WithoutOutputWithArgsChannels struct {
	err chan error
}

func (c WithoutOutputWithArgsChannels) Read() error {
	if len(c.err) == 0 {
		return errors.New("no elements in channel")
	}
	return <-c.err
}

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
