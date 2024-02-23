package parallelize

import (
	"context"
)

type outputArgMethod[I, O any] struct {
	method func(context.Context, I, O) error
	args   OutputArgMethodArgs[I, O]
}

type OutputArgMethodArgs[I, O any] struct {
	Context context.Context
	Arg1    I
	Arg2    O
}

func newOutputArgMethod[I, O any](method func(context.Context, I, O) error, args OutputArgMethodArgs[I, O]) outputArgMethod[I, O] {
	return outputArgMethod[I, O]{
		method: method,
		args:   args,
	}
}

func (m outputArgMethod[I, O]) Execute() error {
	return m.method(m.args.Context, m.args.Arg1, m.args.Arg2)
}
