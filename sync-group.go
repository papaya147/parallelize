package parallelize

import (
	"context"
	"sync"
)

type SyncGroup struct {
	methods []Executable
	errors  *GroupError
}

func NewSyncGroup() *SyncGroup {
	return &SyncGroup{}
}

func AddMethod[I, O any](group *SyncGroup, method func(context.Context, I, O) error, args OutputArgMethodArgs[I, O]) {
	group.methods = append(group.methods, newOutputArgMethod(method, args))
}

func (group *SyncGroup) Run() error {
	group.errors = newGroupError()

	var wg sync.WaitGroup
	errorChannel := make([]chan error, len(group.methods))

	for i, method := range group.methods {
		wg.Add(1)
		errorChannel[i] = make(chan error, 1)
		go func(signature Executable, errorChannel chan error) {
			defer wg.Done()
			errorChannel <- signature.Execute()
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
