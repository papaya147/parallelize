package parallelize

import "context"

// OutputtingMethodWithArgsSignature is a function signature that takes a context and an output pointer and returns an error.
type OutputtingMethodWithoutArgsSignature[O any] func(context.Context, O) error

type outputtingMethodWithoutArgs[O any] struct {
	method OutputtingMethodWithoutArgsSignature[O]
	args   OutputtingMethodWithoutArgsParams[O]
}

// OutputtingMethodWithoutArgsParams is a struct that contains the context and input for a method with arguments.
// The arguments are used by a corresponding OutputtingMethodWithoutArgsSignature
type OutputtingMethodWithoutArgsParams[O any] struct {
	Context context.Context
	Output  O
}

func newOutputtingMethodWithoutArgs[O any](method OutputtingMethodWithoutArgsSignature[O], args OutputtingMethodWithoutArgsParams[O]) executable {
	return outputtingMethodWithoutArgs[O]{
		method: method,
		args:   args,
	}
}

func (m outputtingMethodWithoutArgs[O]) execute() error {
	return m.method(m.args.Context, m.args.Output)
}
