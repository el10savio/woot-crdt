package woot

import "fmt"

// Operation ...
type Operation struct {
	Type      string
	Character WCharacter
}

var (
	// TODO: Convert To String & Tie To IP Address
	SiteID     = 0
	LocalClock = 0
)

// TODO: Unit Tests
// TODO: Bubble up errors

// IsExecutable ...
func (operation *Operation) IsExecutable(wstring WString) bool {
	character := operation.Character

	if operation.Type == "delete" {
		return wstring.Contains(character.ID)
	}

	return wstring.Contains(character.WCPrevious) && wstring.Contains(character.WCNext)
}

// GenerateInsert ...
func (wstring *WString) GenerateInsert(position int, alphabet string) error {
	LocalClock++

	WCharacterPrevious := IthVisible(*wstring, position)
	WCharacterNext := IthVisible(*wstring, position+1)

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
func (wstring *WString) GenerateDelete(position int) {
	wcharacter := IthVisible(*wstring, position)
	wstring.IntegrateDelete(wcharacter)
	// Broadcast
}

// IntegrateDelete ...
func (wstring *WString) IntegrateDelete(wcharacter WCharacter) {
	position, _ := wstring.Position(wcharacter.ID)

	if position == -1 {
		return
	}

	wstring.Sequence[position].Visible = false
}

// IntegrateInsert ...
func (wstring *WString) IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext WCharacter) error {
	var err error

	subsequence, _ := wstring.Subseq(WCharacterPrevious, WCharacterNext)
	position, _ := wstring.Position(WCharacterNext.ID)

	if len(subsequence) == 0 {
		wstring, err = wstring.LocalInsert(wcharacter, position)
		return err
	}

	index := 1

	for index < len(subsequence)-1 && subsequence[index].ID < wcharacter.ID {
		index++
	}

	return wstring.IntegrateInsert(wcharacter, subsequence[index-1], subsequence[index])
}
