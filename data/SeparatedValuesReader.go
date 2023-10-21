package data

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// SeparatedValuesReader is a more generic version of csv.Reader.
type SeparatedValuesReader struct {
	*csv.Reader
}

// NewSeparatedValuesReader returns a new SeparatedValues reader that reads from
// r while using separator and fieldsPerRecord to parse records.
func NewSeparatedValuesReader(r io.Reader, separator rune, fieldsPerRecord int) *SeparatedValuesReader {
	csvReader := csv.NewReader(r)
	csvReader.Comma = separator
	csvReader.FieldsPerRecord = fieldsPerRecord
	csvReader.TrimLeadingSpace = true

	svr := SeparatedValuesReader{
		Reader: csvReader,
	}
	return &svr
}

// ReadSeparatedValues is a utility function for reading values from file
// separated by separator with fieldsPerRecord fields.
func ReadSeparatedValues(file string, separator rune, fieldsPerRecord int) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	defer f.Close()

	r := NewSeparatedValuesReader(f, separator, fieldsPerRecord)
	return r.ReadAll()
}

// ReadSeparatedValuesFloat64 is a utility function for reading values from file
// separated by separator with fieldsPerRecord fields.
func ReadSeparatedValuesFloat64(file string, separator rune, fieldsPerRecord int, skipFirst bool) ([][]float64, error) {
	f, err := os.Open(file)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	defer f.Close()

	r := NewSeparatedValuesReader(f, separator, fieldsPerRecord)

	var values [][]float64

	isFirst := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return values, err
		}
		if skipFirst && isFirst {
			isFirst = false
			continue
		}
		vals := make([]float64, len(record))
		for i := range record {
			v, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return values, err
			}
			vals[i] = v
		}
		values = append(values, vals)
		isFirst = false
	}
	return values, nil
}
