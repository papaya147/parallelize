package parallelize

import "context"

// MethodWithoutArgsSignature is a function signature that takes a context and returns an error.
type MethodWithoutArgsSignature func(context.Context) error

type methodWithoutArgs struct {
	method MethodWithoutArgsSignature
	ctx    context.Context
}

func newMethodWithoutArgs(method MethodWithoutArgsSignature, ctx context.Context) executable {
	return methodWithoutArgs{
		method: method,
		ctx:    ctx,
	}
}

func (m methodWithoutArgs) execute() error {
	return m.method(m.ctx)
}
