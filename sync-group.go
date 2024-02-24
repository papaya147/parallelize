package parallelize

import (
	"context"
	"sync"
)

// SyncGroup is a group of methods that can be executed in parallel.
// All the SyncGroup methods are executed in parallel once the Run method is invoked.
// Add methods to the SyncGroup using the parallelize.AddMethod method.
type SyncGroup struct {
	methods []executable
	errors  *GroupError
}

// NewSyncGroup creates a new SyncGroup.
func NewSyncGroup() *SyncGroup {
	return &SyncGroup{}
}

// AddOutputtingMethoWithArgs adds a method to the SyncGroup.
// The method should take an input and output pointer and return an error.
// The input and output can be of any type.
func AddOutputtingMethodWithArgs[I, O any](group *SyncGroup, method OutputtingMethodWithArgsSignature[I, O], args OutputtingMethodWithArgsParams[I, O]) {
	group.methods = append(group.methods, newOutputtingMethodWithArgs(method, args))
}

// AddOutputtingMethodWithoutArgs adds a method to the SyncGroup.
// The method should take an output pointer and return an error.
// The output can be of any type.
func AddOutputtingMethodWithoutArgs[O any](group *SyncGroup, method OutputtingMethodWithoutArgsSignature[O], args OutputtingMethodWithoutArgsParams[O]) {
	group.methods = append(group.methods, newOutputtingMethodWithoutArgs(method, args))
}

// AddMethodWithArgs adds a method to the SyncGroup.
// The method should take an input and return an error.
// The input can be of any type.
func AddMethodWithArgs[I any](group *SyncGroup, method MethodWithArgsSignature[I], args MethodWithArgsParams[I]) {
	group.methods = append(group.methods, newMethodWithArgs(method, args))
}

// AddMethodWithoutArgs adds a method to the SyncGroup.
// The method should return an error.
func AddMethodWithoutArgs(group *SyncGroup, method MethodWithoutArgsSignature, ctx context.Context) {
	group.methods = append(group.methods, newMethodWithoutArgs(method, ctx))
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
