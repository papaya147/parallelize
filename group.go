package parallelize

import (
	"context"
	"sync"
)

// Group is a group of executables to execute in parallel.
type Group struct {
	executables []executable
}

// NewGroup creates a new Group.
func NewGroup() *Group {
	return &Group{
		executables: []executable{},
	}
}

// AddWithOutputWithArgs adds a new executable to the group with output and with arguments.
func AddWithOutputWithArgs[I, O any](group *Group, f func(context.Context, I) (O, error), ctx context.Context, input I) WithOutputWithArgsChannels[O] {
	m, c := newWithOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

// AddWithOutputWithoutArgs adds a new executable to the group with output and without arguments.
func AddWithOutputWithoutArgs[O any](group *Group, f func(context.Context) (O, error), ctx context.Context) WithOutputWithoutArgsChannels[O] {
	m, c := newWithOutputWithoutArgs(f, ctx)
	group.executables = append(group.executables, m)
	return c
}

// AddWithoutOutputWithArgs adds a new executable to the group without output and with arguments.
func AddWithoutOutputWithArgs[I any](group *Group, f func(context.Context, I) error, ctx context.Context, input I) WithoutOutputWithArgsChannels {
	m, c := newWithoutOutputWithArgs(f, ctx, input)
	group.executables = append(group.executables, m)
	return c
}

// AddWithoutOutputWithoutArgs adds a new executable to the group without output and without arguments.
func AddWithoutOutputWithoutArgs(group *Group, f func(context.Context) error, ctx context.Context) WithoutOutputWithoutArgsChannels {
	m, c := newWithoutOutputWithoutArgs(f, ctx)
	group.executables = append(group.executables, m)
	return c
}

// Execute executes the group of executables in parallel.
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
