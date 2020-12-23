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
	SiteID     = 0 // TODO: Convert To String & Tie To IP Address
	LocalClock = 0

	// TODO: Initialize wstring to contain WCharacterStart & WCharacterEnd
	// TODO: Update WCharacterStart & WCharacterEnd for every wstring operation

	WCharacterStart = WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil}
	WCharacterEnd   = WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil}
)

// Pool is a local var slice of type []Operation{}

// Initialize ...
func Initialize() WString {
	return WString{}
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

// GenerateInsert ...
func (wstring *WString) GenerateInsert(position int, alphabet string) {
	LocalClock++

	WCharacterPrevious := IthVisible(*wstring, position)
	WCharacterNext := IthVisible(*wstring, position+1)

	wcharacter := WCharacter{
		ID:         string(SiteID) + string(LocalClock),
		Visible:    true,
		Alphabet:   alphabet,
		WCPrevious: &WCharacterPrevious,
		WCNext:     &WCharacterNext,
	}

	// IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext)
	// Broadcast
}

// GenerateDelete ...
func (wstring *WString) GenerateDelete(position int) {
	wcharacter := IthVisible(*wstring, position)
	// wstring.IntegrateDelete(wcharacter)
	// Broadcast
}

// Operation ...
type Operation struct {
	Type      string
	Character WCharacter
}

// IsExecutable ...
func (operation *Operation) IsExecutable(wstring WString) bool {
	character := operation.Character

	if operation.Type == "delete" {
		return wstring.Contains(character)
	}

	return wstring.Contains(*character.WCPrevious) && wstring.Contains(*character.WCNext)
}

// IntegrateDelete ...
func (wstring *WString) IntegrateDelete(wcharacter WCharacter) {
	// TODO: Check if position == -1
	position := wstring.Position(wcharacter.ID)
	wstring.Sequence[position].Visible = false
}
