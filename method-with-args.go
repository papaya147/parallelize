package parallelize

import (
	"context"
)

// MethodWithArgsSignature is a function signature that takes a context and an input and returns an error.
type MethodWithArgsSignature[I any] func(context.Context, I) error

type methodWithArgs[I any] struct {
	method MethodWithArgsSignature[I]
	args   MethodWithArgsParams[I]
}

// MethodWithArgsParams is a struct that contains the context and input for a method with arguments.
// The arguments are used by a corresponding MethodWithArgsSignature
type MethodWithArgsParams[I any] struct {
	Context context.Context
	Input   I
}

func newMethodWithArgs[I any](method MethodWithArgsSignature[I], args MethodWithArgsParams[I]) executable {
	return methodWithArgs[I]{
		method: method,
		args:   args,
	}
}

func (m methodWithArgs[O]) execute() error {
	return m.method(m.args.Context, m.args.Input)
}
