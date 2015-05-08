package effpygo

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type FoundWord struct {
	Index int
	Word  string
}

func isLetter(letter rune) bool {
	return !(unicode.IsSpace(letter) || unicode.IsPunct(letter))
}

func IndexWords(text string) []FoundIndex {
	result := make([]FoundIndex, 0)
	var found FoundWord
	var buf bytes.Buffer
	for i, rune := range text {
		if isLetter(rune) {
			if buf.Len() == 0 {
				// First character of a new word
				found.Index = i
			}
			if _, err := buf.WriteRune(r); err != nil {
				panic(err)
			}
		} else {
			if buf.Len() > 0 {
				// Current word is done
				found.Word = buf.String()
				result = append(result, found)
			}
		}
	}
	if buf.Len() > 0 {
		// Output any leftover runes as a word
		found.Word = buf.String()
		result = append(result, found)
	}
	return result
}

type FoundWordOrErr struct {
	FoundWord
	Err error
}

func IndexWordsFromReaderIntoChannel(text io.Reader) <-chan FoundWordOrErr {
	result := make(chan FoundWordOrErr)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				result <- FoundWordOrErr{-1, "", err}
			}
		}()
		defer close(result)

		reader := bufio.NewReader(text)
		var buf bytes.Buffer
		var found FoundWord

		for i := 0; ; i++ {
			rune, _, err := reader.ReadRune()
			if err == io.EOF {
				// Output any remaining runes as the last word
				if buf.Len() > 0 {
					found.Word = buf.String()
					result <- found
				}
				return
			} else if err != nil {
				panic(err)
			}

			if isLetter(rune) {
				if buf.Len() == 0 {
					found.Index = i
				}
				if _, err = buf.WriteRune(rune); err != nil {
					panic(err)
				}
			} else {
				if buf.Len() > 0 {
					found.Word = buf.String()
					result <- found
					buf.Truncate(0)
				}
			}
		}
	}()
	return result
}
