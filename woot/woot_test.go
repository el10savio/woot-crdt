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
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

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
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedWCharacter := WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil}
	actualWCharacter := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}

func Test_ElementAt_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedWCharacter := WCharacter{}
	actualWCharacter := wstring.ElementAt(0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}

func Test_ElementAt_PositionOutOfBounds(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedWCharacter := WCharacter{}
	actualWCharacter := wstring.ElementAt(2)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}

func Test_Postion(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedPosition := 2
	actualPosition := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

func Test_Postion_EmptyWString(t *testing.T) {
	wstring = WString{Sequence: []WCharacter{}}

	expectedPosition := -1
	actualPosition := wstring.Position("end")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

func Test_Postion_EmptyWCharacterID(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedPosition := -1
	actualPosition := wstring.Position("")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

func Test_Postion_WCharacterNotPresent(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedPosition := -1
	actualPosition := wstring.Position("not_present")

	assert.Equal(t, expectedPosition, actualPosition)

	wstring = Clear()
}

// TODO: Test_LocalInsert_Begining
// TODO: Test_LocalInsert_Ending

func Test_LocalInsert(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}
	actualWString := wstring.LocalInsert(wcharacter, 1)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_LocalInsert_PositionOutOfBounds(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedWString := &wstring
	actualWString := wstring.LocalInsert(wcharacter, 3)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_LocalInsert_EmptyWCharacter(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacter := WCharacter{ID: "", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedWString := &wstring
	actualWString := wstring.LocalInsert(wcharacter, 1)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_Subseq(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacterStart := WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: nil, WCNext: nil}
	wcharacterEnd := WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil}

	expectedSubseq := []WCharacter{
		WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
	}
	actualSubseq := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)

	wstring = Clear()
}

func Test_Subseq_WCharacterNotPresent(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacterStart := WCharacter{ID: "a4", Visible: true, Alphabet: "l", WCPrevious: nil, WCNext: nil}
	wcharacterEnd := WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil}

	expectedSubseq := []WCharacter{
		WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: nil, WCNext: nil},
		WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
		WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil},
		WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
	}
	actualSubseq := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)

	wstring = Clear()
}

func Test_Subseq_SameWCharacterRange(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a1", Visible: true, Alphabet: "c", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a3", Visible: true, Alphabet: "t", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacterStart := WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}
	wcharacterEnd := WCharacter{ID: "a2", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedSubseq := []WCharacter{}
	actualSubseq := wstring.Subseq(wcharacterStart, wcharacterEnd)

	assert.Equal(t, expectedSubseq, actualSubseq)

	wstring = Clear()
}

func Test_Contains(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedContains := true
	actualContains := wstring.Contains(wcharacter)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

func Test_Contains_WCharacterNotPresent(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	wcharacter := WCharacter{ID: "a2", Visible: true, Alphabet: "b", WCPrevious: nil, WCNext: nil}

	expectedContains := false
	actualContains := wstring.Contains(wcharacter)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

func Test_Contains_EmptyWString(t *testing.T) {
	wstring = WString{}

	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}

	expectedContains := false
	actualContains := wstring.Contains(wcharacter)

	assert.Equal(t, expectedContains, actualContains)

	wstring = Clear()
}

func Test_Value(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
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
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
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
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedWCharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: nil, WCNext: nil}
	actualWCharacter := IthVisible(wstring, 0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}

func Test_IthVisible_NoVisibleWCharacters(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: nil, WCNext: nil},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: nil, WCNext: nil},
		},
	}

	expectedWCharacter := WCharacter{}
	actualWCharacter := IthVisible(wstring, 0)

	assert.Equal(t, expectedWCharacter, actualWCharacter)

	wstring = Clear()
}
