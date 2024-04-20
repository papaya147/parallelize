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

func AddWithOutputWithArgs[I, O any](group *Group, f func(context.Context, I) (O, error), ctx context.Context, input I) WithOutputWithArgsChannels[O] {
	m, c := newWithOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

func AddWithOutputWithoutArgs[O any](group *Group, f func(context.Context) (O, error), ctx context.Context) WithOutputWithoutArgsChannels[O] {
	m, c := newWithOutputWithoutArgs(f, ctx)
	group.executables = append(group.executables, m)
	return c
}

func AddWithoutOutputWithArgs[I any](group *Group, f func(context.Context, I) error, ctx context.Context, input I) WithoutOutputWithArgsChannels {
	m, c := newWithoutOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

func AddWithoutOutputWithoutArgs(group *Group, f func(context.Context) error, ctx context.Context) WithoutOutputWithoutArgsChannels {
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
