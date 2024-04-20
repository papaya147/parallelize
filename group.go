package parallelize

import (
	"context"
	"sync"
)

type Group struct {
	executables []executable
}

func NewGroup() *Group {
	return &Group{
		executables: []executable{},
	}
}

func NewWithOutputWithArgs[I, O any](group *Group, f func(context.Context, I) (O, error), ctx context.Context, input I) WithOutputWithArgsChannels[O] {
	m, c := newWithOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

func NewWithOutputWithoutArgs[O any](group *Group, f func(context.Context) (O, error), ctx context.Context) WithOutputWithoutArgsChannels[O] {
	m, c := newWithOutputWithoutArgs(f, ctx)
	group.executables = append(group.executables, m)
	return c
}

func NewWithoutOutputWithArgs[I any](group *Group, f func(context.Context, I) error, ctx context.Context, input I) WithoutOutputWithArgsChannels {
	m, c := newWithoutOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

func NewWithoutOutputWithoutArgs(group *Group, f func(context.Context) error, ctx context.Context) WithoutOutputWithoutArgsChannels {
	m, c := newWithoutOutputWithoutArgs(f, ctx)
	group.executables = append(group.executables, m)
	return c
}

func (g *Group) Execute() {
	var wg sync.WaitGroup
	for _, e := range g.executables {
		wg.Add(1)
		go func(e executable) {
			defer wg.Done()
			e.execute()
		}(e)
	}
	wg.Wait()
}
