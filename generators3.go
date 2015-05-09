package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func recordToPoint(record []string) Point {
	var p Point
	if len(record) == 2 {
		p.X, _ = strconv.ParseFloat(record[0], 64)
		p.Y, _ = strconv.ParseFloat(record[1], 64)
	}
	return p
}

func LoadCsvData(text string) (result []Point) {
	reader := csv.NewReader(strings.NewReader(text))
	records, _ := reader.ReadAll()
	for _, record := range records {
		point := recordToPoint(record)
		result = append(result, point)
	}
	return
}

// ---

func LoadCsvDataToChannel(in io.Reader) <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		reader := csv.NewReader(in)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				return
			}
			point := recordToPoint(record)
			out <- point
		}
	}()
	return out
}

// ---

func main() {
	data := "1.0,2.5\n3.5,4.1\n"

	// All at once example
	points := LoadCsvData(data)
	for _, point := range points {
		fmt.Println(point)
	}

	// Streaming example
	results := LoadCsvDataToChannel(strings.NewReader(data))
	for point := range results {
		fmt.Println(point)
	}
}
