package woot

import "fmt"

var (
	// TODO: Convert SiteID To String & Tie To IP Address

	// SiteID is the global unique variable
	// used to identify the WOOT node site
	SiteID = 0

	// LocalClock is a local clock incremented whenever
	// an insert operation takes place to
	// uniquely identify each WCharacter
	LocalClock = 0
)

// Find returns the WCharacter present
// in the WString Sequence for
// the given WCharacter ID
func (wstring *WString) Find(ID string) WCharacter {
	// Iterate over the WString Sequence and
	// search for the WCharacter based
	// on its WCharacter ID
	for _, wcharacter := range wstring.Sequence {
		if wcharacter.ID == ID {
			return wcharacter
		}
	}

	// Return a WCharacter with ID -1
	// When no WCharacter is found
	return WCharacter{ID: "-1"}
}

// GenerateInsert generates a WCharacter for a given alphabet
// into the WString Sequence at the given position
func (wstring *WString) GenerateInsert(position int, alphabet string) (*WString, error) {
	// Increment the LocalClock
	LocalClock++

	// Obtain the previous WCharacter
	WCharacterPrevious := IthVisible(*wstring, position-1)

	// Obtain the next WCharacter
	WCharacterNext := IthVisible(*wstring, position)

	// Default to WCharacterStart if
	// WCharacterPrevious is not found
	if WCharacterPrevious.ID == "-1" {
		WCharacterPrevious = wstring.Find("start")
	}

	// Default to WCharacterEnd if
	// WCharacterNext is not found
	if WCharacterNext.ID == "-1" {
		WCharacterNext = wstring.Find("end")
	}

	// Generate the WCharacter
	wcharacter := WCharacter{
		ID:         fmt.Sprint(SiteID) + fmt.Sprint(LocalClock),
		Visible:    true,
		Alphabet:   alphabet,
		WCPrevious: WCharacterPrevious.ID,
		WCNext:     WCharacterNext.ID,
	}

	return wstring.IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext)
}

// GenerateDelete generates the WCharacter to be marked for
// deletion in the WString Sequence at the given position
func (wstring *WString) GenerateDelete(position int) *WString {
	wcharacter := IthVisible(*wstring, position)
	return wstring.IntegrateDelete(wcharacter)
}

// IntegrateDelete finds out the position of the given WCharacter
// in the WString Sequence and marks it for deletion
// by setting the visible flag for it to false
func (wstring *WString) IntegrateDelete(wcharacter WCharacter) *WString {
	// Obtain the position of the given WCharacter
	position := wstring.Position(wcharacter.ID)

	// Guard checks for the position to check if it's in the WString
	// bounds If it is out of bounds the same WString is returned and
	// prevents the WString Sequence indexing to panic
	if position == -1 {
		return wstring
	}

	// Mark the given element as visible = false
	wstring.Sequence[position-1].Visible = false

	// Return the updated WString
	return wstring
}

// IntegrateInsert inserts the given WCharacter into the WString
// Sequence based off of the previous & next WCharacter
func (wstring *WString) IntegrateInsert(wcharacter, WCharacterPrevious, WCharacterNext WCharacter) (*WString, error) {
	// Get the subsequence present between the previous & next WCharacter
	subsequence, _ := wstring.Subseq(WCharacterPrevious, WCharacterNext)

	// Get the position of the next WCharacter
	position := wstring.Position(WCharacterNext.ID)

	// Since it's a natural number
	// decrement it by 1
	position--

	// If no WCharacters are present in the subseqence
	// LocalInsert it at the given position
	if len(subsequence) == 0 {
		return wstring.LocalInsert(wcharacter, position)
	}

	// If a single WCharacter is present in the subseqence
	// LocalInsert it at the previous position
	if len(subsequence) == 1 {
		return wstring.LocalInsert(wcharacter, position-1)
	}

	// Recuresively figure out the correct bounds
	// for the previous & next WCharacter

	index := 1

	for index < len(subsequence)-1 && subsequence[index].ID < wcharacter.ID {
		index++
	}

	return wstring.IntegrateInsert(wcharacter, subsequence[index-1], subsequence[index])
}
