// Copyright 2015 Brett Slatkin, Pearson Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func recordToPoint(record []string) (p Point) {
	if len(record) != 2 {
		return
	}
	p.X, _ = strconv.ParseFloat(record[0], 64)
	p.Y, _ = strconv.ParseFloat(record[1], 64)
	return
}

func LoadCsvData(in io.Reader) (result []Point) {
	reader := csv.NewReader(in)
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
	points := LoadCsvData(strings.NewReader(data))
	for i, point := range points {
		fmt.Printf("Row %d is %v\n", i, point)
	}

	// Streaming example
	results := LoadCsvDataToChannel(strings.NewReader(data))
	i := 0
	for point := range results {
		fmt.Printf("Row %d is %v\n", i, point)
		i++
	}
}
