package woot

import "errors"

// WString ...
type WString struct {
	Sequence []WCharacter // Ordered sequence of WCharacters
}

// WCharacter ...
type WCharacter struct {
	ID         string // Identifier of the character
	Visible    bool   // Is the character visible
	Alphabet   string // Alphabetical value of the effect character
	WCPrevious string // Identifier of the previous WCharacter
	WCNext     string // Identifier of the next WCharacter
}

var (
	// WCharacterStart ...
	WCharacterStart = WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "end"}

	// WCharacterEnd ...
	WCharacterEnd = WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "start", WCNext: ""}

	// TODO: Move errors to its own file

	ErrPositionOutOfBounds = errors.New("position out of bounds")
	ErrEmptyWCharacter     = errors.New("empty wcharacter ID provided")
	ErrBoundsNotPresent    = errors.New("subsequence bound(s) not present")
)

// Pool is a local var slice of type []Operation{}

// Initialize ...
func Initialize() WString {
	return WString{Sequence: []WCharacter{WCharacterStart, WCharacterEnd}}
}

// Length ...
func (wstring *WString) Length() int {
	return len(wstring.Sequence)
}

// ElementAt ...
func (wstring *WString) ElementAt(position int) (WCharacter, error) {
	if position < 0 || position >= wstring.Length() {
		return WCharacter{}, ErrPositionOutOfBounds
	}
	return wstring.Sequence[position], nil
}

// Position ...
// Returns wstring natural number
func (wstring *WString) Position(wcharacterID string) (int, error) {
	if wcharacterID == "" {
		return -1, ErrEmptyWCharacter
	}

	for position, _wcharacter := range wstring.Sequence {
		if wcharacterID == _wcharacter.ID {
			return position + 1, nil
		}
	}

	return -1, nil

}

// LocalInsert ...
func (wstring *WString) LocalInsert(wcharacter WCharacter, position int) (*WString, error) {
	if position <= 0 || position >= wstring.Length() {
		return wstring, ErrPositionOutOfBounds
	}

	if wcharacter.ID == "" {
		return wstring, ErrEmptyWCharacter
	}

	wstring.Sequence = append(wstring.Sequence[:position],
		append([]WCharacter{wcharacter}, wstring.Sequence[position:]...)...,
	)

	wstring.Sequence[position].WCPrevious = wstring.Sequence[position-1].ID
	wstring.Sequence[position].WCNext = wstring.Sequence[position+1].ID

	wstring.Sequence[position-1].WCNext = wcharacter.ID
	wstring.Sequence[position+1].WCPrevious = wcharacter.ID

	return wstring, nil
}

// Subseq ...
// Excluding wcharacterStart & wcharacterEnd
func (wstring *WString) Subseq(wcharacterStart, wcharacterEnd WCharacter) ([]WCharacter, error) {
	startPosition, _ := wstring.Position(wcharacterStart.ID)
	endPosition, _ := wstring.Position(wcharacterEnd.ID)

	if startPosition == -1 || endPosition == -1 {
		return wstring.Sequence, ErrBoundsNotPresent
	}

	// Same WCharacter Start & End
	if startPosition == endPosition {
		return []WCharacter{}, nil
	}

	return wstring.Sequence[startPosition : endPosition-1], nil
}

// Contains ...
func (wstring *WString) Contains(wcharacterID string) bool {
	position, _ := wstring.Position(wcharacterID)
	return position != -1
}

// Value ...
func Value(wstring WString) string {
	value := ""

	for _, wcharacter := range wstring.Sequence {
		if wcharacter.Visible {
			value += wcharacter.Alphabet
		}
	}

	return value
}

// IthVisible ...
func IthVisible(wstring WString, position int) WCharacter {
	count := 0

	for _, wcharacter := range wstring.Sequence {
		if wcharacter.Visible {
			if count == position {
				return wcharacter
			}
			count++
		}
	}

	return WCharacter{}
}
