package generators2

import (
	"csv"
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
}

func LoadCsvData(text string) (result []Point, err error) {
	reader := csv.NewReader(strings.NewReader(text))
	records, err := reader.ReadAll()
	if err != nil {
		return
	}
	for _, record := range records {
		point, err := recordToPoint(record)
		if err != nil {
			return
		}
		result = append(result, point)
	}
	return
}

type recordOrErr struct {
	record []string
	err    error
}

func loadCsvUntilEof(in io.Reader) <-chan recordOrErr {
	out := make(chan recordOrErr)
	go func() {
		defer close(out)
		reader := csv.NewReader(in)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				out <- recordOrErr{err: err}
				return
			}
			out <- recordOrErr{record: record}
		}
	}()
	return out
}

type PointOrErr struct {
	Point
	Err error
}

func LoadCsvDataToChannel(in io.Reader) <-chan PointOrErr {
	out := make(chan Point)
	go func() {
		defer close(out)
		for record := range loadCsvUntilEof(in) {
			if record.err != nil {
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
