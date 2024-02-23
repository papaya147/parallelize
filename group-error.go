package parallelize

import (
	"fmt"
)

type GroupError struct {
	Errors []error
}

func newGroupError() *GroupError {
	return &GroupError{Errors: []error{}}
}

func (e *GroupError) Add(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *GroupError) Error() string {
	errStr := ""
	for _, err := range e.Errors {
		if err == nil {
			continue
		}
		errStr = fmt.Sprintf("%s, %s", errStr, err.Error())
	}
	if errStr == "" {
		return errStr
	}
	return errStr[2:]
}

func (e *GroupError) IsEmpty() bool {
	return len(e.Errors) == 0
}
