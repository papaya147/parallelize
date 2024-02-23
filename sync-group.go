package parallelize

import (
	"context"
	"sync"
)

// SyncGroup is a group of methods that can be executed in parallel.
type SyncGroup struct {
	methods []executable
	errors  *GroupError
}

// NewSyncGroup creates a new SyncGroup.
func NewSyncGroup() *SyncGroup {
	return &SyncGroup{}
}

// AddMethod adds a method to the SyncGroup.
// The method should take an input and output and return an error.
// The input and output can be any type.
func AddMethod[I, O any](group *SyncGroup, method func(context.Context, I, O) error, args OutputArgMethodArgs[I, O]) {
	group.methods = append(group.methods, newOutputArgMethod(method, args))
}

// Run executes the methods in the SyncGroup in parallel.
// It returns an error if any of the methods return an error.
func (group *SyncGroup) Run() error {
	group.errors = newGroupError()

	var wg sync.WaitGroup
	errorChannel := make([]chan error, len(group.methods))

	for i, method := range group.methods {
		wg.Add(1)
		errorChannel[i] = make(chan error, 1)
		go func(method executable, errorChannel chan error) {
			defer wg.Done()
			errorChannel <- method.execute()
		}(method, errorChannel[i])
	}

	wg.Wait()

	for _, ch := range errorChannel {
		if err := <-ch; err != nil {
			group.errors.Add(err)
		}
	}

	if group.errors.IsEmpty() {
		return nil
	}
	return group.errors
}
