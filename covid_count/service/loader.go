package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Loader struct {
	Data *repository.Manager
}

func NewLoader(rm *repository.Manager) Loader{
	return Loader{
		Data: rm,
	}
}

func (l *Loader) Load() error{
	path := l.getLatestFile()
	csvReader := GetReaderForCSV(path)

	lineCount := l.getLineCount(path)
	err := l.Data.IsUpdating().SetIsUpdating(true)
	if err != nil {
		return err
	}
	err = l.Data.Record().ClearPreviousRecords()
	if err != nil {
		return err
	}
	records := l.csvToRecords(csvReader, lineCount)
	err = l.Data.Record().CreateMultiple(records)
	if err != nil {
		return err
	}
	return l.Data.IsUpdating().SetIsUpdating(false)
}


func (l *Loader) csvToRecords(csvReader *csv.Reader, lineCount int) []domain.Record {
	var records []domain.Record
	currentLineCount := 0
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		// Skip the first and last line
		currentLineCount++
		if currentLineCount == 1 {
			continue
		}
		if currentLineCount == lineCount {
			continue
		}
		record, err := domain.NewRecordFromCSV(line)
		records = append(records, record)
	}
	return records
}


func (l *Loader) getLineCount(path string) int {
	lineReader, _ := os.Open(path)
	lineCount, _ := lineCounter(lineReader)
	return lineCount
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}



func (l *Loader) getLatestFile() string {
	summaryURL := "https://coronavirus.ohio.gov/static/COVIDSummaryData.csv"
	path := "./data/imports/" + currentDate() + ".csv"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		l.downloadFile(path, summaryURL)
	}
	return path
}

func currentDate() string {
	now := time.Now()
	return fmt.Sprintf("%d%d%d", now.Day(), now.Month(), now.Year())
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (l *Loader) downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
