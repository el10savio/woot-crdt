package woot

import "fmt"

// Operation ...
type Operation struct {
	Type      string
	Character WCharacter
}

var (
	SiteID     = 0 // TODO: Convert To String & Tie To IP Address
	LocalClock = 0
)

// IsExecutable ...
func (operation *Operation) IsExecutable(wstring WString) bool {
	character := operation.Character

	if operation.Type == "delete" {
		return wstring.Contains(character)
	}

	return wstring.Contains(*character.WCPrevious) && wstring.Contains(*character.WCNext)
}

// GenerateInsert ...
func (wstring *WString) GenerateInsert(position int, alphabet string) {
	LocalClock++

	WCharacterPrevious := IthVisible(*wstring, position)
	WCharacterNext := IthVisible(*wstring, position+1)

	wcharacter := WCharacter{
		ID:         fmt.Sprint(SiteID) + fmt.Sprint(LocalClock),
		Visible:    true,
		Alphabet:   alphabet,
		WCPrevious: &WCharacterPrevious,
		WCNext:     &WCharacterNext,
	}

	wstring.IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext)

	// Broadcast
}

// GenerateDelete ...
func (wstring *WString) GenerateDelete(position int) {
	wcharacter := IthVisible(*wstring, position)
	wstring.IntegrateDelete(wcharacter)
	// Broadcast
}

// IntegrateDelete ...
func (wstring *WString) IntegrateDelete(wcharacter WCharacter) {
	// TODO: Check if position == -1
	position := wstring.Position(wcharacter.ID)
	wstring.Sequence[position].Visible = false
}

// IntegrateInsert ...
func (wstring *WString) IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext WCharacter) {
	subsequence := wstring.Subseq(WCharacterPrevious, WCharacterNext)

	if len(subsequence) == 0 {
		wstring = wstring.LocalInsert(wcharacter, wstring.Position(WCharacterNext.ID))
		return
	}

	index := 1

	for index < len(subsequence)-1 && subsequence[index].ID < wcharacter.ID {
		index++
	}

	wstring.IntegrateInsert(wcharacter, subsequence[index-1], subsequence[index])
}
