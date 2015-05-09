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

func recordToPoint(record []string) (p Point, err error) {
	if len(record) != 2 {
		err = fmt.Errorf("Records must have two columns")
		return
	}
	p.X, err = strconv.ParseFloat(record[0], 64)
	if err != nil {
		return
	}
	p.Y, err = strconv.ParseFloat(record[1], 64)
	if err != nil {
		return
	}
	return
}

func LoadCsvData(in io.Reader) ([]Point, error) {
	var result []Point
	reader := csv.NewReader(in)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		point, err := recordToPoint(record)
		if err != nil {
			return nil, err
		}
		result = append(result, point)
	}
	return result, nil
}

// ---

type PointOrErr struct {
	Point
	Err error
}

func LoadCsvDataToChannel(in io.Reader) <-chan PointOrErr {
	out := make(chan PointOrErr)
	go func() {
		defer close(out)
		reader := csv.NewReader(in)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				out <- PointOrErr{Err: err}
				return
			}
			point, err := recordToPoint(record)
			if err != nil {
				out <- PointOrErr{Err: err}
				return
			}
			out <- PointOrErr{Point: point}
		}
	}()
	return out
}

// ---

func main() {
	data := "1.0,2.5\n3.5,4.1\n"

	// All at once example
	points, err := LoadCsvData(strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	for _, point := range points {
		fmt.Println(point)
	}

	// Streaming example
	results := LoadCsvDataToChannel(strings.NewReader(data))
	for point := range results {
		if point.Err != nil {
			panic(point.Err)
		}
		fmt.Println(point.Point)
	}
}
