package woot

import "errors"

var (
	// ErrPositionOutOfBounds is the error message sent
	// when the position out of bounds
	ErrPositionOutOfBounds = errors.New("position out of bounds")

	// ErrEmptyWCharacter is the error message sent
	// when an empty wcharacter is ID provided
	ErrEmptyWCharacter = errors.New("empty wcharacter ID provided")

	// ErrBoundsNotPresent is the error message sent
	// when subsequence bound(s) are not present
	ErrBoundsNotPresent = errors.New("subsequence bound(s) not present")
)
