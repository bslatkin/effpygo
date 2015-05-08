package effpygo

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type Word struct {
	Index int
	Text  string
}

func isLetter(r rune) bool {
	return !(unicode.IsSpace(r) || unicode.IsPunct(r))
}

func readUntil(targetStatus bool, reader *bufio.Reader) string {
	var buf bytes.Buffer
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if isLetter(r) == targetStatus {
			reader.UnreadRune()
			break
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func IndexWords(in io.Reader) []Word {
	words := make([]Word, 0)
	reader := bufio.NewReader(in)
	index := 0
	for {
		whitespace := readUntil(true, reader)
		index += len(whitespace)
		text := readUntil(false, reader)
		if len(text) == 0 {
			break
		}
		words = append(words, Word{index, text})
		index += len(text)
	}
	return words
}

type WordOrErr struct {
	Word
	Err error
}

func IndexWordsStream(in io.Reader) <-chan Word {
	out := make(chan Word)
	go func() {
		defer close(out)
		reader := bufio.NewReader(in)
		index := 0
		for {
			whitespace := readUntil(true, reader)
			index += len(whitespace)
			text := readUntil(false, reader)
			if len(text) == 0 {
				break
			}
			out <- Word{index, text}
			index += len(text)
		}
	}()
	return out
}
