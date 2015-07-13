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

package effpygo

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
)

type Point struct {
	X, Y float64
}

func recordToPoint(record []string) (p Point, err error) {
	if len(record) != 2 {
		err = fmt.Errorf("Records must have two columns")
		return
	}
	if p.X, err = strconv.ParseFloat(record[0], 64); err != nil {
		return
	}
	if p.Y, err = strconv.ParseFloat(record[1], 64); err != nil {
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

type DistanceOrErr struct {
	Distance float64
	Err      error
}

func PointDistanceToChannel(in <-chan PointOrErr) <-chan DistanceOrErr {
	out := make(chan DistanceOrErr)
	go func() {
		defer close(out)
		p := <-in
		if p.Err != nil {
			out <- DistanceOrErr{Err: p.Err}
		}
		for q := range in {
			if q.Err != nil {
				out <- DistanceOrErr{Err: q.Err}
				continue
			}
			dx := math.Pow(q.X-p.X, 2)
			dy := math.Pow(q.Y-p.Y, 2)
			distance := math.Sqrt(dx + dy)
			out <- DistanceOrErr{Distance: distance}
			p = q
		}
	}()
	return out
}
