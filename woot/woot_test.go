package woot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// wstring is the local WString
	// variable stored for tests
	wstring WString
)

func init() {
	// Initialize the
	// wstring variable
	wstring = Initialize()
}

// Clear is a utility function used
// to reset the wstring variable
// at the end of tests
func Clear() WString {
	return Initialize()
}

// Test_Length tests the basic functionality
// of the wstring Length() function
func Test_Length(t *testing.T) {
	wstring = Initialize()

	expectedLength := 2
	actualLength := wstring.Length()

	assert.Equal(t, expectedLength, actualLength)

	wstring = Clear()
}

// Test_Length_Empty tests the basic functionality
// of the wstring Length() function
// when the wstring is empty
func Test_Length_Empty(t *testing.T) {
	wstring = WString{}

	expectedLength := 0
	actualLength := wstring.Length()

	assert.Equal(t, expectedLength, actualLength)

	wstring = Clear()
}

// Test_ElementAt tests the basic functionality
// of the wstring ElementAt() function
func Test_ElementAt(t *testing.T) {
	wstring = Initialize()

	expectedWCharacter := WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID}
	actualWCharacter, actualErr := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_ElementAt_EmptyWString tests the basic functionality
// of the wstring ElementAt() function
// when the wstring is empty
func Test_ElementAt_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedWCharacter := WCharacter{}
	expectedErr := ErrPositionOutOfBounds

	actualWCharacter, actualErr := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_ElementAt_PositionOutOfBounds tests the basic functionality
// of the wstring ElementAt() function when the position
// is out of bounds of the wstring sequence
func Test_ElementAt_PositionOutOfBounds(t *testing.T) {
	wstring = Initialize()

	expectedWCharacter := WCharacter{}
	expectedErr := ErrPositionOutOfBounds

	actualWCharacter, actualErr := wstring.ElementAt(2)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_Position tests the basic functionality
// of the wstring Positihon() function
func Test_Position(t *testing.T) {
	wstring = Initialize()

	expectedPosition := 2
	actualPosition := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

// Test_Position_EmptyWString tests the basic functionality
// of the wstring Position() function
// when the wstring is empty
func Test_Position_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedPosition := -1
	actualPosition := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

// Test_Position_EmptyWCharacterID tests the basic functionality
// of the wstring Position() function
// when the WCharacter ID is empty
func Test_Position_EmptyWCharacterID(t *testing.T) {
	wstring = Initialize()

	expectedPosition := -1

	actualPosition := wstring.Position("")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

// Test_Position_EmptyWCharacterID tests the basic functionality
// of the wstring Position() function
// when the WCharacter is not present
func Test_Position_WCharacterNotPresent(t *testing.T) {
	wstring = Initialize()

	expectedPosition := -1
	actualPosition := wstring.Position("not_present")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

// Test_LocalInsert tests the basic functionality
// of the wstring LocalInsert() function
func Test_LocalInsert(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}
	actualWString, actualErr := wstring.LocalInsert(wcharacter, 1)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_Beginning tests the basic functionality
// of the wstring LocalInsert() function when inserting
// to the beginning of the wstring sequence
func Test_LocalInsert_Beginning(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedWString := &wstring
	expectedErr := ErrPositionOutOfBounds

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 0)

	assert.Equal(t, expectedWString, actualWString)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_Ending tests the basic functionality
// of the wstring LocalInsert() function when inserting
// to the ending of the wstring sequence
func Test_LocalInsert_Ending(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedWString := &wstring
	expectedErr := ErrPositionOutOfBounds

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_ReplaceBegining tests the basic functionality
// of the wstring LocalInsert() function when replacing
// the beginning of the wstring sequence
func Test_LocalInsert_ReplaceBegining(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "a", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "b", Visible: true, Alphabet: "b", WCPrevious: "start", WCNext: "a"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "b"},
			WCharacter{ID: "b", Visible: true, Alphabet: "b", WCPrevious: "start", WCNext: "a"},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "b", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "a", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 1)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_ReplaceEnding tests the basic functionality
// of the wstring LocalInsert() function when replacing
// the ending of the wstring sequence
func Test_LocalInsert_ReplaceEnding(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "a", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "b", Visible: true, Alphabet: "b", WCPrevious: "a", WCNext: "end"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "b"},
			WCharacter{ID: "b", Visible: true, Alphabet: "b", WCPrevious: "a", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_PositionOutOfBounds tests the basic functionality
// of the wstring LocalInsert() function when the position
// is out of bounds of the wstring sequence
func Test_LocalInsert_PositionOutOfBounds(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedWString := &wstring
	expectedErr := ErrPositionOutOfBounds

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 3)

	assert.Equal(t, expectedWString, actualWString)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_LocalInsert_EmptyWCharacter tests the basic functionality
// of the wstring LocalInsert() function when the WCharacter is empty
func Test_LocalInsert_EmptyWCharacter(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedWString := &wstring
	expectedErr := ErrEmptyWCharacter
	actualWString, actualErr := wstring.LocalInsert(wcharacter, 1)

	assert.Equal(t, expectedWString, actualWString)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_Subseq tests the basic functionality
// of the wstring Subseq() function
func Test_Subseq(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	wcharacterStart := WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: "", WCNext: ""}
	wcharacterEnd := WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""}

	expectedSubseq := []WCharacter{
		WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
	}
	actualSubseq, actualErr := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_Subseq_WCharacterNotPresent tests the basic functionality
// of the wstring Subseq() function when the WCharacter
// bounds are not present in the wstring sequence
func Test_Subseq_WCharacterNotPresent(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	wcharacterStart := WCharacter{ID: "a4", Visible: true, Alphabet: "l", WCPrevious: "", WCNext: ""}
	wcharacterEnd := WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""}

	expectedSubseq := []WCharacter{
		WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: "", WCNext: ""},
		WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
		WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""},
		WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
	}
	expectedErr := ErrBoundsNotPresent

	actualSubseq, actualErr := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

// Test_Subseq_SameWCharacterRange tests the basic functionality
// of the wstring Subseq() function when the WCharacter bounds are
// not presentpoint to the same WCharacter in the wstring sequence
func Test_Subseq_SameWCharacterRange(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	wcharacterStart := WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}
	wcharacterEnd := WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedSubseq := []WCharacter{}
	actualSubseq, actualErr := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

// Test_Contains tests the basic functionality
// of the wstring Contains() function
func Test_Contains(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedContains := true
	actualContains := wstring.Contains(wcharacter.ID)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

// Test_Contains_WCharacterNotPresent tests the basic functionality
// of the wstring Contains() function when the WCharacter
// is not present in the wstring
func Test_Contains_WCharacterNotPresent(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "a2", Visible: true, Alphabet: "b", WCPrevious: "", WCNext: ""}

	expectedContains := false
	actualContains := wstring.Contains(wcharacter.ID)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

// Test_Contains_EmptyWString tests the basic functionality
// of the wstring Contains() function when the WCharacter is empty
func Test_Contains_EmptyWString(t *testing.T) {
	wstring = WString{}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedContains := false
	actualContains := wstring.Contains(wcharacter.ID)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

// Test_Value tests the basic functionality
// of the Value() function
func Test_Value(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	expectedValue := "a"
	actualValue := Value(wstring)

	assert.Equal(t, expectedValue, actualValue)

	wstring = Clear()
}

// Test_Value_NoVisibleWCharacters tests the basic functionality
// of the Value() function when no visible WCharacters
// are present in the wstring
func Test_Value_NoVisibleWCharacters(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	expectedValue := ""
	actualValue := Value(wstring)

	assert.Equal(t, expectedValue, actualValue)

	wstring = Clear()
}

// Test_IthVisible tests the basic functionality of the IthVisible() function
func Test_IthVisible(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	expectedWCharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}
	actualWCharacter := IthVisible(wstring, 1)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}

// Test_IthVisible_NoVisibleWCharacters tests the basic functionality
// of the IthVisible() function when no visible WCharacters
// are present in the wstring
func Test_IthVisible_NoVisibleWCharacters(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: "", WCNext: ""},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""},
		},
	}

	expectedWCharacter := WCharacter{ID: "-1", Visible: false, Alphabet: "", WCPrevious: "", WCNext: ""}
	actualWCharacter := IthVisible(wstring, 2)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}
