package util

import (
	"errors"
	"strings"
	"sync"
)

// ErrorCollector collects errors.
type ErrorCollector struct {
	errs   []error
	errMap map[string]bool

	lock sync.RWMutex
}

func (e *ErrorCollector) Append(err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	_, exists := e.errMap[err.Error()]
	if !exists {
		e.errs = append(e.errs, err)
		e.errMap[err.Error()] = true
	}
}

func (e *ErrorCollector) Error() error {
	e.lock.RLock()
	defer e.lock.RUnlock()

	switch e.Length() {
	case 0:
		return errors.New("no error")
	case 1:
		return e.errs[0]
	default:
		var sb strings.Builder
		sb.WriteString("multiple errors: ")

		lastErr := e.Length() - 1
		for i, err := range e.errs {
			sb.WriteString(err.Error())
			if i != lastErr {
				sb.WriteString(", ")
			}
		}

		return errors.New(sb.String())
	}
}

func (e *ErrorCollector) Length() int {
	e.lock.RLock()
	defer e.lock.RUnlock()

	return len(e.errs)
}
