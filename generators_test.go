package effpygo

import (
	"fmt"
	"strings"
)

const (
	Address           = "Four score and seven years ago our fathers brought forth on this continent a new nation, conceived in liberty, and dedicated to the proposition that all men are created equal."
	AddressWithSpaces = "  Four score and seven years ago our fathers brought forth on this   continent a new nation, conceived in liberty, and dedicated to the proposition that all men are created equal.  "
)

func ExampleIndexWords() {
	result := IndexWords(Address)
	fmt.Println(result)
	// Output:
	// [0 5 11 15 21 27 31 35 43 51 57 60 65 75 77 81 89 99 102 111 115 125 128 132 144 149 153 157 161 169]
}

func ExampleIndexWordsWithSpaces() {
	text := AddressWithSpaces
	indexes := IndexWords(text)
	for _, found := range indexes {
		word := text[found.Index : found.Index+found.Length]
		fmt.Println(found.Index, word)
	}
	// Output:
	// 2 Four
	// 7 score
	// 13 and
	// 17 seven
	// 23 years
	// 29 ago
	// 33 our
	// 37 fathers
	// 45 brought
	// 53 forth
	// 59 on
	// 62 this
	// 69 continent
	// 79 a
	// 81 new
	// 85 nation
	// 93 conceived
	// 103 in
	// 106 liberty
	// 115 and
	// 119 dedicated
	// 129 to
	// 132 the
	// 136 proposition
	// 148 that
	// 153 all
	// 157 men
	// 161 are
	// 165 created
	// 173 equal
}

func ExampleIndexWordsIntoChannel() {
	for index := range IndexWordsIntoChannel(Address) {
		fmt.Println(index)
	}
	// Output:
	// 0
	// 5
	// 11
	// 15
	// 21
	// 27
	// 31
	// 35
	// 43
	// 51
	// 57
	// 60
	// 65
	// 75
	// 77
	// 81
	// 89
	// 99
	// 102
	// 111
	// 115
	// 125
	// 128
	// 132
	// 144
	// 149
	// 153
	// 157
	// 161
	// 169
}

func ExampleIndexWordsFromReaderIntoChannel() {
	input := strings.NewReader(Address)
	for found := range IndexWordsFromReaderIntoChannel(input) {
		if found.Err != nil {
			panic(found.Err)
		}
		fmt.Println(found)
	}
	// Output:
	// asdf
}
