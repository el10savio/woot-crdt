package woot

// WString is used to define our text sequence
// It is a collection of WCharacters and
// Is used to save the state of the text sequence
type WString struct {
	Sequence []WCharacter // Ordered sequence of WCharacters
}

// WCharacter is a struct used to store
// information about a character in our text
type WCharacter struct {
	ID         string // Identifier of the character
	Visible    bool   // Is the character visible
	Alphabet   string // Alphabetical value of the effect character
	WCPrevious string // Identifier of the previous WCharacter
	WCNext     string // Identifier of the next WCharacter
}

var (
	// WCharacterStart is a special WCharacter placed at the start
	// of the WString to denote the start of the WCharacter Sequence
	WCharacterStart = WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "end"}

	// WCharacterEnd is a special WCharacter placed at the end
	// of the WString to denote the end of the WCharacter Sequence
	WCharacterEnd = WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "start", WCNext: ""}
)

// Initialize returns a WString initialized
// with WCharacterStart & WCharacterEnd
func Initialize() WString {
	return WString{Sequence: []WCharacter{WCharacterStart, WCharacterEnd}}
}

// Length returns the length
// of the WString Sequence
func (wstring *WString) Length() int {
	return len(wstring.Sequence)
}

// ElementAt returns the WCharacter
// present in the given position
// in the WString Sequence
func (wstring *WString) ElementAt(position int) (WCharacter, error) {
	// Guard checks for position to check if it's in the WString
	// bounds If it is out of bounds an error is returned and
	// prevents the WString Sequence indexing to panic
	if position < 0 || position >= wstring.Length() {
		return WCharacter{}, ErrPositionOutOfBounds
	}

	// Return WCharacter at the given position
	// in the WString Sequence
	return wstring.Sequence[position], nil
}

// Position returns the position of the given WCharacter
// in the WString Sequence as a natural number
func (wstring *WString) Position(wcharacterID string) int {
	// Iterate over the WString Sequence and match
	// the WCharacter by its given WCharacter ID
	for position, _wcharacter := range wstring.Sequence {
		if wcharacterID == _wcharacter.ID {
			return position + 1
		}
	}

	// Return -1 if the given WCharacter
	// is not present
	return -1
}

// LocalInsert inserts the given WCharacter into the WString Sequence at the given
// position and shifts the remaining elements to accommodate the new WCharacter
func (wstring *WString) LocalInsert(wcharacter WCharacter, position int) (*WString, error) {
	// Guard checks for position to check if it's in the WString
	// bounds If it is out of bounds an error is returned and
	// prevents the WString Sequence indexing to panic
	if position <= 0 || position >= wstring.Length() {
		return wstring, ErrPositionOutOfBounds
	}

	// Guard checks for WCharacter to check if it's ID is empty
	// If it has an empty ID an error is returned
	if wcharacter.ID == "" {
		return wstring, ErrEmptyWCharacter
	}

	// Add the new WCharacter into the WString Sequence
	// and shift the remaining elements
	wstring.Sequence = append(wstring.Sequence[:position],
		append([]WCharacter{wcharacter}, wstring.Sequence[position:]...)...,
	)

	// Update the WCNext field of the previous WCharacter
	wstring.Sequence[position-1].WCNext = wcharacter.ID

	// Update the WCPrevious field of the next WCharacter
	wstring.Sequence[position+1].WCPrevious = wcharacter.ID

	// Return the updated WString
	return wstring, nil
}

// Subseq return the part of the WString Sequence
// between wcharacterStart & wcharacterEnd
func (wstring *WString) Subseq(wcharacterStart, wcharacterEnd WCharacter) ([]WCharacter, error) {
	// Get the position of wcharacterStart & wcharacterEnd
	startPosition := wstring.Position(wcharacterStart.ID)
	endPosition := wstring.Position(wcharacterEnd.ID)

	// Guard checks for positions to check if they are in the WString
	// If any one is not present an error is returned and
	// prevents the WString Sequence indexing to panic
	if startPosition == -1 || endPosition == -1 {
		return wstring.Sequence, ErrBoundsNotPresent
	}

	// Guard checks for positions to check if both is pointing to the
	// same WCharacter present then an empty WCharacter is returned
	// and prevents the WString Sequence indexing to panic
	if startPosition == endPosition {
		return []WCharacter{}, nil
	}

	// Return the Subsequence between the
	// startPosition & endPosition
	return wstring.Sequence[startPosition : endPosition-1], nil
}

// Contains returns a boolean if the given
// WCharacter is present in the WString
func (wstring *WString) Contains(wcharacterID string) bool {
	position := wstring.Position(wcharacterID)
	return position != -1
}

// Value is the representation of all the
// visible characters in the WString
func Value(wstring WString) string {
	// Initialize the value
	// to an empty string
	value := ""

	// Iterate over the visible WCharacters
	// and add them to the value variable
	for _, wcharacter := range wstring.Sequence {
		if wcharacter.Visible {
			value += wcharacter.Alphabet
		}
	}

	// Return the combined value
	return value
}

// IthVisible returns the ith visible
// WCharacter in the WString
func IthVisible(wstring WString, position int) WCharacter {
	// Initialize a marker update when
	// a visible WCharacter is found
	count := 0

	// Iterate over the visible WCharacters
	// and update the count, when the count is
	// equal to the position return the WCharacter
	for _, wcharacter := range wstring.Sequence {
		if wcharacter.Visible {
			if count == position-1 {
				return wcharacter
			}
			count++
		}
	}

	// Return a WCharacter with ID -1
	// When no WCharacter is found
	return WCharacter{ID: "-1"}
}
