package service

import (
	"archive/zip"
	"fmt"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type DeathLoader struct {
	Data *repository.Manager
}

func NewDeathLoader(data *repository.Manager) *DeathLoader {
	return &DeathLoader{Data: data}
}

func (dl *DeathLoader) Load() error{
	path := "./data/imports/" + currentDate() + "deaths.zip"
	err := dl.downloadDeathFile(path)
	if err != nil {
		return err
	}
	//err = dl.Data.DeathRecords().ClearPreviousMonthlyCountyDeathss()
	//if err != nil {
	//	return err
	//}
	//var records []domain.MonthlyCountyDeaths
	//for currentYear := 2018; currentYear <= 2020; currentYear++ {
	//	csvReader := GetReaderForCSV("./data/imports/"+strconv.Itoa(currentYear)+"deaths.csv")
	//	for {
	//		line, err := csvReader.Read()
	//		if err == io.EOF {
	//			break
	//		}
	//
	//		if err != nil {
	//			return err
	//		}
	//
	//		record, err := domain.NewMonthlyCountyDeathsFromCSV(line, currentYear)
	//		records = append(records, record...)
	//	}
	//}
	//
	//err = dl.Data.DeathRecords().CreateMultiple(records)
	//if err != nil{
	//	return err
	//}

	return nil
}



// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (dl *DeathLoader) downloadDeathFile(filepath string) error {

	client := &http.Client{

	}
	req, _ := http.NewRequest("GET", "http://publicapps.odh.ohio.gov/EDW/Reports/GetDelimitedReportData", nil)
	req.Header.Add("Cookie", "ASP.NET_SessionId=japbgqfpw1vf3i0l35osx1tg; ai_user=j9yht|2020-10-17T18:38:57.482Z; ai_session=KKnNB|1602959939593|1602960176385.54")
	resp, err := client.Do(req)
	// Get the data
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
	files, err := Unzip(filepath, "./data/imports/")
	fmt.Println(files)
	return err
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
