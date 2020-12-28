package woot

import "fmt"

// Operation ...
type Operation struct {
	Type      string
	Character WCharacter
}

var (
	// TODO: Convert SiteID To String & Tie To IP Address

	// SiteID ...
	SiteID = 0

	// LocalClock ...
	LocalClock = 0
)

// IsExecutable ...
func (operation *Operation) IsExecutable(wstring WString) bool {
	character := operation.Character

	if operation.Type == "delete" {
		return wstring.Contains(character.ID)
	}

	return wstring.Contains(character.WCPrevious) && wstring.Contains(character.WCNext)
}

// Find ...
func (wstring *WString) Find(ID string) WCharacter {
	for _, wcharacter := range wstring.Sequence {
		if wcharacter.ID == ID {
			return wcharacter
		}
	}
	return WCharacter{ID: "-1"}
}

// GenerateInsert ...
func (wstring *WString) GenerateInsert(position int, alphabet string) (*WString, error) {
	LocalClock++

	WCharacterPrevious := IthVisible(*wstring, position-1)
	WCharacterNext := IthVisible(*wstring, position)

	if WCharacterPrevious.ID == "-1" {
		WCharacterPrevious = wstring.Find("start")
	}

	if WCharacterNext.ID == "-1" {
		WCharacterNext = wstring.Find("end")
	}

	wcharacter := WCharacter{
		ID:         fmt.Sprint(SiteID) + fmt.Sprint(LocalClock),
		Visible:    true,
		Alphabet:   alphabet,
		WCPrevious: WCharacterPrevious.ID,
		WCNext:     WCharacterNext.ID,
	}

	return wstring.IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext)

	// Broadcast
}

// GenerateDelete ...
func (wstring *WString) GenerateDelete(position int) *WString {
	wcharacter := IthVisible(*wstring, position)
	return wstring.IntegrateDelete(wcharacter)
	// Broadcast
}

// IntegrateDelete ...
func (wstring *WString) IntegrateDelete(wcharacter WCharacter) *WString {
	position := wstring.Position(wcharacter.ID)

	if position == -1 {
		return wstring
	}

	wstring.Sequence[position-1].Visible = false

	return wstring
}

// IntegrateInsert ...
func (wstring *WString) IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext WCharacter) (*WString, error) {

	subsequence, _ := wstring.Subseq(WCharacterPrevious, WCharacterNext)
	position := wstring.Position(WCharacterNext.ID)

	position--

	if len(subsequence) == 0 {
		return wstring.LocalInsert(wcharacter, position)
	}

	if len(subsequence) == 1 {
		return wstring.LocalInsert(wcharacter, position-1)
	}

	index := 1

	for index < len(subsequence)-1 && subsequence[index].ID < wcharacter.ID {
		index++
	}

	return wstring.IntegrateInsert(wcharacter, subsequence[index-1], subsequence[index])
}
