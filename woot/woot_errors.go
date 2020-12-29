package woot

import "errors"

var (
	ErrPositionOutOfBounds = errors.New("position out of bounds")
	ErrEmptyWCharacter     = errors.New("empty wcharacter ID provided")
	ErrBoundsNotPresent    = errors.New("subsequence bound(s) not present")
)
