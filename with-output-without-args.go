package parallelize

import (
	"context"
	"errors"
)

type WithOutputWithoutArgsSignature[O any] func(context.Context) (O, error)

type WithOutputWithoutArgsChannels[O any] struct {
	output chan O
	err    chan error
}

func (c WithOutputWithoutArgsChannels[O]) Read() (O, error) {
	if len(c.output) == 0 || len(c.err) == 0 {
		return *new(O), errors.New("no elements in channel, maybe you didn't execute?")
	}
	return <-c.output, <-c.err
}

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
