package woot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	wstring WString
)

func init() {
	wstring = Initialize()
}

func Clear() WString {
	return Initialize()
}

func Test_Length(t *testing.T) {
	wstring = Initialize()

	expectedLength := 2
	actualLength := wstring.Length()

	assert.Equal(t, expectedLength, actualLength)

	wstring = Clear()
}

func Test_Length_Empty(t *testing.T) {
	wstring = WString{}

	expectedLength := 0
	actualLength := wstring.Length()

	assert.Equal(t, expectedLength, actualLength)

	wstring = Clear()
}

func Test_ElementAt(t *testing.T) {
	wstring = Initialize()

	expectedWCharacter := WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID}
	actualWCharacter, actualErr := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_ElementAt_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedWCharacter := WCharacter{}
	expectedErr := ErrPositionOutOfBounds

	actualWCharacter, actualErr := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

func Test_ElementAt_PositionOutOfBounds(t *testing.T) {
	wstring = Initialize()

	expectedWCharacter := WCharacter{}
	expectedErr := ErrPositionOutOfBounds

	actualWCharacter, actualErr := wstring.ElementAt(2)

	assert.Equal(t, expectedWCharacter, actualWCharacter)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

func Test_Postion(t *testing.T) {
	wstring = Initialize()

	expectedPosition := 2
	actualPosition, actualErr := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_Postion_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedPosition := -1
	actualPosition, actualErr := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_Postion_EmptyWCharacterID(t *testing.T) {
	wstring = Initialize()

	expectedPosition := -1
	expectedErr := ErrEmptyWCharacter

	actualPosition, actualErr := wstring.Position("")

	assert.Equal(t, expectedPosition, actualPosition)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

func Test_Postion_WCharacterNotPresent(t *testing.T) {
	wstring = Initialize()

	expectedPosition := -1
	actualPosition, actualErr := wstring.Position("not_present")

	assert.Equal(t, expectedPosition, actualPosition)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

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

func Test_LocalInsert_Begining(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedWString := &wstring
	expectedErr := ErrPositionOutOfBounds

	actualWString, actualErr := wstring.LocalInsert(wcharacter, 0)

	assert.Equal(t, expectedWString, actualWString)
	assert.Equal(t, expectedErr, actualErr)

	wstring = Clear()
}

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

func Test_Contains_EmptyWString(t *testing.T) {
	wstring = WString{}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: "", WCNext: ""}

	expectedContains := false
	actualContains := wstring.Contains(wcharacter.ID)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

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
