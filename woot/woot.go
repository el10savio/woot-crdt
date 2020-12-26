package woot

// WString ...
type WString struct {
	Sequence []WCharacter // Ordered sequence of WCharacters
}

// WCharacter ...
type WCharacter struct {
	ID       string // Identifier of the character
	Visible  bool   // Is the character visible
	Alphabet string // Alphabetical value of the effect character

	// TODO: Change Previous & Next WCharacter to store ID

	WCPrevious *WCharacter // Identifier of the previous WCharacter
	WCNext     *WCharacter // Identifier of the next WCharacter
}

var (
	// TODO: Update WCharacterStart & WCharacterEnd for every wstring operation

	// WCharacterStart ...
	WCharacterStart = WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil}

	// WCharacterEnd ...
	WCharacterEnd = WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil}
)

// Pool is a local var slice of type []Operation{}

// Initialize ...
// TODO: Initialize wstring to contain WCharacterStart & WCharacterEnd
func Initialize() WString {
	_WCharacterStart, _WCharacterEnd := WCharacterStart, WCharacterEnd
	_WCharacterStart.WCNext, _WCharacterEnd.WCPrevious = &WCharacterEnd, &WCharacterStart
	return WString{Sequence: []WCharacter{_WCharacterStart, _WCharacterEnd}}
}

// Length ...
func (wstring *WString) Length() int {
	return len(wstring.Sequence)
}

// ElementAt ...
func (wstring *WString) ElementAt(position int) WCharacter {
	if position < 0 || position >= wstring.Length() {
		// TODO: Return err
		return WCharacter{}
	}
	return wstring.Sequence[position]
}

// Position ...
// Returns wstring natural number
func (wstring *WString) Position(wcharacterID string) int {
	if wcharacterID == "" {
		// TODO: Return err
		return -1
	}

	for position, _wcharacter := range wstring.Sequence {
		if wcharacterID == _wcharacter.ID {
			return position + 1
		}
	}

	return -1
}

// LocalInsert ...
func (wstring *WString) LocalInsert(wcharacter WCharacter, position int) *WString {
	if position < 0 || position >= wstring.Length() {
		// TODO: Return err
		return wstring
	}

	if wcharacter.ID == "" {
		// TODO: Return err
		return wstring
	}

	wstring.Sequence = append(wstring.Sequence[:position],
		append([]WCharacter{wcharacter}, wstring.Sequence[position:]...)...,
	)

	return wstring
}

// Subseq ...
// Excluding wcharacterStart & wcharacterEnd
func (wstring *WString) Subseq(wcharacterStart, wcharacterEnd WCharacter) []WCharacter {
	startPosition := wstring.Position(wcharacterStart.ID)
	endPosition := wstring.Position(wcharacterEnd.ID)

	// TODO: Return err if position == -1
	if startPosition == -1 || endPosition == -1 {
		return wstring.Sequence
	}

	// Same WCharacter Start & End
	if startPosition == endPosition {
		return []WCharacter{}
	}

	return wstring.Sequence[startPosition : endPosition-1]
}

// Contains ...
func (wstring *WString) Contains(wcharacter WCharacter) bool {
	position := wstring.Position(wcharacter.ID)
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
