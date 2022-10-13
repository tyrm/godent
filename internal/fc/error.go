package fc

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrNotOKStatus        = errors.New("got non 200 status response")
	ErrHomeServerNotFound = errors.New("home server not found")
)
