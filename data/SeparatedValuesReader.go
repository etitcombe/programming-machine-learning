package data

import (
	"encoding/csv"
	"io"
	"os"
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
