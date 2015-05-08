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
	for _, found := range IndexWords(Address) {
		fmt.Println(found)
	}
	// Output:
	// {0 4}
	// {5 5}
	// {11 3}
	// {15 5}
	// {21 5}
	// {27 3}
	// {31 3}
	// {35 7}
	// {43 7}
	// {51 5}
	// {57 2}
	// {60 4}
	// {65 9}
	// {75 1}
	// {77 3}
	// {81 6}
	// {89 9}
	// {99 2}
	// {102 7}
	// {111 3}
	// {115 9}
	// {125 2}
	// {128 3}
	// {132 11}
	// {144 4}
	// {149 3}
	// {153 3}
	// {157 3}
	// {161 7}
	// {169 5}
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
