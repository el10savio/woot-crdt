package woot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntegrateInsert(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(wcharacter, WCharacterStart, WCharacterEnd)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateInsert_Middle(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: "a", WCNext: "b"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "x"},
			WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: "a", WCNext: "b"},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "x", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(
		wcharacter,
		WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
		WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
	)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateInsert_Transpose(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: WCharacterStart.ID, WCNext: "a"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "x"},
			WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: WCharacterStart.ID, WCNext: "a"},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: "x", WCNext: "b"},
			WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(
		wcharacter,
		WCharacterStart,
		WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
	)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateDelete(t *testing.T) {
	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	actualWString := wstring.IntegrateDelete(wcharacter)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_IntegrateDelete_WCharacterNotPresent(t *testing.T) {
	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: WCharacterStart.ID, WCNext: ""},
		},
	}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: WCharacterStart.ID, WCNext: ""},
		},
	}

	actualWString := wstring.IntegrateDelete(wcharacter)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_GenerateInsert(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position, alphabet := 1, "a"

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.GenerateInsert(position, alphabet)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_ReplaceStart(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position := 1
	alphabet1, alphabet2 := "a", "b"

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "02"},
			WCharacter{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "start", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "02", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(position, alphabet1)
	actualWString, actualErr := WStringPtr.GenerateInsert(position, alphabet2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_ReplaceEnd(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabet1, alphabet2 := "a", "b"

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "02"},
			WCharacter{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "01", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "02", WCNext: ""},
		},
	}

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, alphabet1)
	actualWString, actualErr := WStringPtr.GenerateInsert(2, alphabet2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}
	LocalClock = 0

	position := 1

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	actualWString := wstring.GenerateDelete(position)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete_NoValue(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position := 1

	expectedWString := &wstring

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(position, "a")
	WStringPtr = WStringPtr.GenerateDelete(position)
	actualWString := WStringPtr.GenerateDelete(position)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
	LocalClock = 0
}
