package effpygo

import (
	"fmt"
	"strings"
)

const (
	Address           = "Four score and seven years ago our fathers brought forth on this continent a new nation, conceived in liberty, and dedicated to the proposition that all men are created equal."
	AddressWithSpaces = "  Four score and seven years ago our fathers brought forth on this   continent a new nation, conceived in liberty, and dedicated to the proposition that all men are created equal.  "
	NoEndingLetter    = "Four    score and    seven"
)

func ExampleIndexWords() {
	for _, word := range IndexWords(Address) {
		fmt.Println(word)
	}
	// Output:
	// {0 Four}
	// {5 score}
	// {11 and}
	// {15 seven}
	// {21 years}
	// {27 ago}
	// {31 our}
	// {35 fathers}
	// {43 brought}
	// {51 forth}
	// {57 on}
	// {60 this}
	// {65 continent}
	// {75 a}
	// {77 new}
	// {81 nation}
	// {89 conceived}
	// {99 in}
	// {102 liberty}
	// {111 and}
	// {115 dedicated}
	// {125 to}
	// {128 the}
	// {132 proposition}
	// {144 that}
	// {149 all}
	// {153 men}
	// {157 are}
	// {161 created}
	// {169 equal}
}

func ExampleIndexWords_WithSpaces() {
	for _, word := range IndexWords(AddressWithSpaces) {
		fmt.Println(word)
	}
	// Output:
	// {2 Four}
	// {7 score}
	// {13 and}
	// {17 seven}
	// {23 years}
	// {29 ago}
	// {33 our}
	// {37 fathers}
	// {45 brought}
	// {53 forth}
	// {59 on}
	// {62 this}
	// {69 continent}
	// {79 a}
	// {81 new}
	// {85 nation}
	// {93 conceived}
	// {103 in}
	// {106 liberty}
	// {115 and}
	// {119 dedicated}
	// {129 to}
	// {132 the}
	// {136 proposition}
	// {148 that}
	// {153 all}
	// {157 men}
	// {161 are}
	// {165 created}
	// {173 equal}
}

func ExampleIndexWords_NoEndingLetter() {
	for _, word := range IndexWords(NoEndingLetter) {
		fmt.Println(word)
	}
	// Output:
	// {0 Four}
	// {8 score}
	// {14 and}
	// {21 seven}
}

func ExampleIndexWordsStream() {
	for word := range IndexWordsStream(strings.NewReader(Address)) {
		fmt.Println(word)
	}
	// Output:
	// {0 Four}
	// {5 score}
	// {11 and}
	// {15 seven}
	// {21 years}
	// {27 ago}
	// {31 our}
	// {35 fathers}
	// {43 brought}
	// {51 forth}
	// {57 on}
	// {60 this}
	// {65 continent}
	// {75 a}
	// {77 new}
	// {81 nation}
	// {89 conceived}
	// {99 in}
	// {102 liberty}
	// {111 and}
	// {115 dedicated}
	// {125 to}
	// {128 the}
	// {132 proposition}
	// {144 that}
	// {149 all}
	// {153 men}
	// {157 are}
	// {161 created}
	// {169 equal}
}

func ExampleIndexWordsStream_WithSpaces() {
	for word := range IndexWordsStream(strings.NewReader(AddressWithSpaces)) {
		fmt.Println(word)
	}
	// Output:
	// {2 Four}
	// {7 score}
	// {13 and}
	// {17 seven}
	// {23 years}
	// {29 ago}
	// {33 our}
	// {37 fathers}
	// {45 brought}
	// {53 forth}
	// {59 on}
	// {62 this}
	// {69 continent}
	// {79 a}
	// {81 new}
	// {85 nation}
	// {93 conceived}
	// {103 in}
	// {106 liberty}
	// {115 and}
	// {119 dedicated}
	// {129 to}
	// {132 the}
	// {136 proposition}
	// {148 that}
	// {153 all}
	// {157 men}
	// {161 are}
	// {165 created}
	// {173 equal}
}
