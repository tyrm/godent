package fc

import "errors"

var (
	ErrNotOKStatus        = errors.New("got non 200 status response")
	ErrHomeServerNotFound = errors.New("home server not found")
)
