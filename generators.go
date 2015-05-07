package effpygo

import (
	"bufio"
	"io"
	"unicode"
)

type FoundIndex struct {
	Index, Length int
}

func isLetter(letter rune) bool {
	return !(unicode.IsSpace(letter) || unicode.IsPunct(letter))
}

func IndexWords(text string) []FoundIndex {
	result := make([]FoundIndex, 0)

	lastIndex := -1
	ignoreLetter := true
	for index, letter := range text {
		if ignoreLetter {
			// Consume ignorable until the first letter of a word
			if isLetter(letter) {
				ignoreLetter = false
				lastIndex = index
			}
		} else {
			// Consume letters from word until the first to ignore
			if !isLetter(letter) {
				ignoreLetter = true
				result = append(result, FoundIndex{lastIndex, index - lastIndex})
			}
		}
	}
	if !ignoreLetter {
		// The last word won't be output unless there is trailing whitespace
		result = append(result, FoundIndex{lastIndex, len(text) - lastIndex})
	}
	return result
}

func IndexWordsIntoChannel(text string) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		if len(text) > 0 {
			result <- 0
		}
		for index, letter := range text {
			if letter == ' ' {
				result <- index + 1
			}
		}
	}()
	return result
}

type FoundWord struct {
	Index int
	Word  string
	Err   error
}

func IndexWordsFromReaderIntoChannel(text io.Reader) <-chan FoundWord {
	result := make(chan FoundWord)
	go func() {
		defer close(result)

		reader := bufio.NewReader(text)
		index := 0
		for {
			next, err := reader.ReadString(' ')
			if err == io.EOF {
				return
			}
			if err != nil {
				result <- FoundWord{-1, "", err}
				return
			}
			trimmed := next[:len(next)-1]
			if len(trimmed) > 0 {
				result <- FoundWord{index, trimmed, nil}
			}
			index += len(next)
		}
	}()
	return result
}
