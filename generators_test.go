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
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	data = "1.0,2.5\n3.5,4.1\n7.5,2.2\n6.9,1.1\n"
)

func ExampleAllAtOnce() {
	points, err := LoadCsvData(strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	for i, point := range points {
		fmt.Printf("Row %d is %v\n", i, point)
	}
	// Output:
	// Row 0 is {1 2.5}
	// Row 1 is {3.5 4.1}
	// Row 2 is {7.5 2.2}
	// Row 3 is {6.9 1.1}
}

func ExampleStreaming() {
	results := LoadCsvDataToChannel(strings.NewReader(data))
	i := 0
	for point := range results {
		if point.Err != nil {
			panic(point.Err)
		}
		fmt.Printf("Row %d is %v\n", i, point)
		i++
	}
	// Output:
	// Row 0 is {{1 2.5} <nil>}
	// Row 1 is {{3.5 4.1} <nil>}
	// Row 2 is {{7.5 2.2} <nil>}
	// Row 3 is {{6.9 1.1} <nil>}
}

func ExampleDistance() {
	pointStream := LoadCsvDataToChannel(strings.NewReader(data))
	distances := PointDistanceToChannel(pointStream)
	i := 0
	for distance := range distances {
		if distance.Err != nil {
			panic(distance.Err)
		}
		fmt.Printf("Move %d was %f far\n", i, distance.Distance)
		i++
	}
	// Output:
	// Move 0 was 2.968164 far
	// Move 1 was 4.428318 far
	// Move 2 was 1.252996 far
}

func TestDistanceBadFirstPoint(t *testing.T) {
	pointStream := make(chan PointOrErr)
	go func() {
		pointStream <- PointOrErr{Err: fmt.Errorf("Bad first point")}
		close(pointStream)
	}()
	distances := PointDistanceToChannel(pointStream)
	found := <-distances
	if found.Err.Error() != "Bad first point" {
		t.Fail()
	}
	if _, open := <-distances; open {
		t.Fail()
	}
}

func TestDistanceBadLaterPoint(t *testing.T) {
	pointStream := make(chan PointOrErr)
	go func() {
		pointStream <- PointOrErr{Point: Point{1.0, 0.0}}
		pointStream <- PointOrErr{Point: Point{4.0, 0.0}}
		pointStream <- PointOrErr{Err: fmt.Errorf("Bad point")}
		pointStream <- PointOrErr{Point: Point{10.0, 0.0}}
		close(pointStream)
	}()
	distances := PointDistanceToChannel(pointStream)

	found := <-distances
	if found.Distance != 3.0 {
		t.Fatalf("First output was %#v", found)
	}

	found = <-distances
	if !(found.Err != nil && found.Err.Error() == "Bad point") {
		t.Fatalf("Second output was %#v", found)
	}

	found = <-distances
	if found.Distance != 6.0 {
		t.Fatalf("Third output was %#v", found)
	}

	if _, open := <-distances; open {
		t.Fail()
	}
}

func getCsvData() string {
	var buf bytes.Buffer
	for i := 0; i < 100000; i++ {
		_, err := buf.WriteString("1.5,2.5\n")
		if err != nil {
			panic(err)
		}
	}
	return buf.String()
}

func BenchmarkAllAtOnce(b *testing.B) {
	data := getCsvData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoadCsvData(strings.NewReader(data))
	}
}

func BenchmarkStreaming(b *testing.B) {
	data := getCsvData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		points := LoadCsvDataToChannel(strings.NewReader(data))
		for p := range points {
			if p.Err != nil {
				panic(p.Err)
			}
		}
	}
}
