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
		for _ = range points {
			// Do nothing
		}
	}
}
