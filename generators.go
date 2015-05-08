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

func isPartOfWord(r rune) bool {
	return !(unicode.IsSpace(r) || unicode.IsPunct(r))
}

// ---

func IndexWords(text string) (result []Word) {
	var word Word

	for i, r := range text {
		if isPartOfWord(r) {
			// When the buffer was empty, but now we've found a
			// rune that's part of a word, mark the index as the
			// first character of a newly found word.
			if len(word.Text) == 0 {
				word.Index = i
			}
			// Always append the rune to the current word.
			word.Text += string(r)
		} else {
			// When the current rune is whitespace or punctuation,
			// then we may have reached the end of the word and
			// need to save a new result.
			if len(word.Text) > 0 {
				result = append(result, word)
				word.Text = ""
			}
		}
	}

	// Any runes remaining in the buffer after we've gone through the text
	// should be returned as part of a final found word.
	if len(word.Text) > 0 {
		result = append(result, word)
		word.Text = ""
	}

	return
}

// ---

func IndexWordsFromStream(reader io.RuneReader) (result []Word, err error) {
	var word Word

	for i := 0; ; i++ {
		var r rune
		if r, _, err = reader.ReadRune(); err != nil {
			if err == io.EOF {
				err = nil
				if len(word.Text) > 0 {
					result = append(result, word)
					word.Text = ""
				}
			}
			return
		}

		if isPartOfWord(r) {
			if len(word.Text) == 0 {
				word.Index = i
			}
			word.Text += string(r)
		} else {
			if len(word.Text) > 0 {
				result = append(result, word)
				word.Text = ""
			}
		}
	}

	return
}

// ---

func readUntil(targetIsPartOfWord bool, reader io.RuneScanner) (string, error) {
	var buf bytes.Buffer
	for {
		var r rune
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if isPartOfWord(r) == targetIsPartOfWord {
			if err := reader.UnreadRune(); err != nil {
				return "", err
			}
			break
		}
		if err := buf.WriteRune(r); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

// ---

func IndexWordsToChannel(text string) <-chan Word {
	result := make(chan Word)
	func() {
		defer close(result)
		var word Word

		for i, r := range text {
			if isPartOfWord(r) {
				if len(word.Text) == 0 {
					word.Index = i
				}
				word.Text += string(r)
			} else {
				if len(word.Text) > 0 {
					result <- word
					word.Text = ""
				}
			}
		}

		if len(word.Text) > 0 {
			result <- word
			word.Text = ""
		}
	}()
	return result
}

// ---

// func getNext() rune {
// 	r, _, err := reader.ReadRune()
// 	if err != nil {
// 		return 0
// 	}
// }

// type statusChange struct {
// 	isText bool
// 	word   string
// }

// func generateStatusChanges(reader *bufio.Reader) <-chan statusChange {
// 	r, _, err := reader.ReadRune()
// 	if err != nil {
// 		return
// 	}
// 	nextStatus := isLetter(r)

// 	for {
// 		r, _, err := reader.ReadRune()
// 		if err != nil {
// 			return
// 		}
// 		next = isLetter(r)
// 		if next != current {

// 		}
// 	}
// }

func readUntil(targetStatus bool, reader *bufio.Reader) string {
	var buf bytes.Buffer
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if isPartOfWord(r) == targetStatus {
			reader.UnreadRune()
			break
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

// func IndexWordsFromReader(in io.Reader) []Word {
// 	words := make([]Word, 0)
// 	reader := bufio.NewReader(in)
// 	index := 0
// 	for {
// 		whitespace := readUntil(true, reader)
// 		index += len(whitespace)
// 		text := readUntil(false, reader)
// 		if len(text) == 0 {
// 			break
// 		}
// 		words = append(words, Word{index, text})
// 		index += len(text)
// 	}
// 	return words
// }

// ---

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
