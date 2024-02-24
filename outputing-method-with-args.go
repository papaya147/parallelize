package parallelize

import (
	"context"
)

// OutputtingMethodWithArgsSignature is a function signature that takes a context, an input
// and an output pointer and returns an error.
type OutputtingMethodWithArgsSignature[I, O any] func(context.Context, I, O) error

type outputtingMethodWithArgs[I, O any] struct {
	method OutputtingMethodWithArgsSignature[I, O]
	args   OutputtingMethodWithArgsParams[I, O]
}

// OutputtingMethodWithArgsParams is a struct that contains the context and input for a method with arguments.
// The arguments are used by a corresponding OutputtingMethodWithArgsSignature
type OutputtingMethodWithArgsParams[I, O any] struct {
	Context context.Context
	Input   I
	Output  O
}

func newOutputtingMethodWithArgs[I, O any](method OutputtingMethodWithArgsSignature[I, O], args OutputtingMethodWithArgsParams[I, O]) executable {
	return outputtingMethodWithArgs[I, O]{
		method: method,
		args:   args,
	}
}

func (m outputtingMethodWithArgs[I, O]) execute() error {
	return m.method(m.args.Context, m.args.Input, m.args.Output)
}
