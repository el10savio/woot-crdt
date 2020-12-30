package woot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_IntegrateInsert tests the basic functionality
// of the IntegrateInsert function
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

// Test_IntegrateInsert_Middle tests the basic functionality
// of the IntegrateInsert function when inserting
// into the middle of the WString Sequence
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

// Test_IntegrateInsert_Transpose tests the basic functionality
// of the IntegrateInsert function when inserting into
// the WString Sequence when a WCharacter is
// already present in the given position
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

// Test_IntegrateDelete tests the basic functionality
// of the IntegrateDelete function
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

// Test_IntegrateDelete_WCharacterNotPresent tests the basic functionality
// of the IntegrateDelete function when the
// WCharacter to be deleted is not present
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

// Test_GenerateInsert tests the basic functionality
// of the GenerateInsert function
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

// Test_GenerateInsert_ReplaceStart tests the basic functionality
// of the GenerateInsert function when replacing the
// WCharacter in the start of thw WString Sequence
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

// Test_GenerateInsert_ReplaceEnd tests the basic functionality
// of the GenerateInsert function when replacing the
// WCharacter in the end of thw WString Sequence
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

// Test_GenerateInsert_Word tests the basic functionality
// of the GenerateInsert function when generating a word
func Test_GenerateInsert_Word(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabets := []string{"a", "b", "c"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "02"},
			WCharacter{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "01", WCNext: "03"},
			WCharacter{ID: "03", Visible: true, Alphabet: "c", WCPrevious: "02", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "03", WCNext: ""},
		},
	}

	var actualErr error
	count := 1

	actualWString, _ := wstring.GenerateInsert(count, alphabets[0])

	for _, alphabet := range alphabets[1:] {
		count++
		actualWString, actualErr = wstring.GenerateInsert(count, alphabet)
	}

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	expectedText := "abc"
	actualText := Value(*actualWString)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

// Test_GenerateInsert_Sentence tests the basic functionality
// of the GenerateInsert function when generating a sentence
func Test_GenerateInsert_Sentence(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabets := []string{"I", " ", "l", "i", "k", "e", " ", "d", "o", "g", "s"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			WCharacter{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			WCharacter{ID: "01", Visible: true, Alphabet: "I", WCPrevious: "start", WCNext: "02"},
			WCharacter{ID: "02", Visible: true, Alphabet: " ", WCPrevious: "01", WCNext: "03"},
			WCharacter{ID: "03", Visible: true, Alphabet: "l", WCPrevious: "02", WCNext: "04"},
			WCharacter{ID: "04", Visible: true, Alphabet: "i", WCPrevious: "03", WCNext: "05"},
			WCharacter{ID: "05", Visible: true, Alphabet: "k", WCPrevious: "04", WCNext: "06"},
			WCharacter{ID: "06", Visible: true, Alphabet: "e", WCPrevious: "05", WCNext: "07"},
			WCharacter{ID: "07", Visible: true, Alphabet: " ", WCPrevious: "06", WCNext: "08"},
			WCharacter{ID: "08", Visible: true, Alphabet: "d", WCPrevious: "07", WCNext: "09"},
			WCharacter{ID: "09", Visible: true, Alphabet: "o", WCPrevious: "08", WCNext: "010"},
			WCharacter{ID: "010", Visible: true, Alphabet: "g", WCPrevious: "09", WCNext: "011"},
			WCharacter{ID: "011", Visible: true, Alphabet: "s", WCPrevious: "010", WCNext: "end"},
			WCharacter{ID: "end", Visible: false, Alphabet: "", WCPrevious: "011", WCNext: ""},
		},
	}

	var actualErr error
	count := 1

	actualWString, _ := wstring.GenerateInsert(count, alphabets[0])

	for _, alphabet := range alphabets[1:] {
		count++
		actualWString, actualErr = wstring.GenerateInsert(count, alphabet)
	}

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	expectedText := "I like dogs"
	actualText := Value(*actualWString)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

// Test_GenerateDelete tests the basic functionality
// of the GenerateDelete function
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

// Test_GenerateDelete tests the basic functionality
// of the GenerateDelete function when
// the value to be deleted is empty
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

// Test_GenerateDelete_WordReplaceInPlace tests the basic functionality
// of the GenerateDelete function when replacing a word
func Test_GenerateDelete_WordReplaceInPlace(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, "1")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "2")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "3")
	WStringPtr = WStringPtr.GenerateDelete(2)
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "4")

	expectedText := "143"
	actualText := Value(*WStringPtr)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

// Test_GenerateDelete_SentenceReplaceInPlace tests the basic functionality
// of the GenerateDelete function when replacing a sentence
func Test_GenerateDelete_SentenceReplaceInPlace(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, "1")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "2")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "3")
	WStringPtr, _ = WStringPtr.GenerateInsert(4, " ")

	WStringPtr, _ = wstring.GenerateInsert(5, "4")
	WStringPtr, _ = WStringPtr.GenerateInsert(6, "5")
	WStringPtr, _ = WStringPtr.GenerateInsert(7, "6")
	WStringPtr, _ = WStringPtr.GenerateInsert(8, " ")

	WStringPtr, _ = wstring.GenerateInsert(9, "7")
	WStringPtr, _ = WStringPtr.GenerateInsert(10, "8")
	WStringPtr, _ = WStringPtr.GenerateInsert(11, "9")
	WStringPtr, _ = WStringPtr.GenerateInsert(12, " ")

	WStringPtr = wstring.GenerateDelete(1)
	WStringPtr = WStringPtr.GenerateDelete(1)
	WStringPtr = WStringPtr.GenerateDelete(1)

	WStringPtr, _ = wstring.GenerateInsert(1, "7")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "8")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "9")

	WStringPtr = wstring.GenerateDelete(9)
	WStringPtr = WStringPtr.GenerateDelete(9)
	WStringPtr = WStringPtr.GenerateDelete(9)

	WStringPtr, _ = wstring.GenerateInsert(9, "1")
	WStringPtr, _ = wstring.GenerateInsert(10, "2")
	WStringPtr, _ = wstring.GenerateInsert(11, "3")

	expectedText := "789 456 123 "
	actualText := Value(*WStringPtr)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}
