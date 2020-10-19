package service

import (
	"encoding/csv"
	"os"
)

func GetReaderForCSV(path string) *csv.Reader {
	reader, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(reader)
	return csvReader
}
