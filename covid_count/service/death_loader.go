package service

import (
	"archive/zip"
	"fmt"
	"github.com/maxheckel/covid_county/covid_count/domain"
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
	files, err := dl.downloadDeathFiles(path)
	if err != nil {
		return err
	}
	err = dl.Data.DeathRecords().ClearPreviousMonthlyCountyDeaths()
	if err != nil {
		return err
	}
	var records []domain.MonthlyCountyDeaths
	for _, file := range files {
		csvReader := GetReaderForCSV(file)
		for {
			line, err := csvReader.Read()

			if err == io.EOF {
				break
			}
			// Skip header values and dummy data
			if line[0] == "DeathMonthMonthName" || line[0] == "Total" || line[0] == "Unknown" || line[1] == "NonOH"{
				continue
			}
			if err != nil {
				return err
			}
			year := 0
			if strings.Contains(file, "2020") {
				year = 2020
			}
			record, err := domain.NewMonthlyCountyDeathsFromCSV(line, year)
			records = append(records, record)
		}
	}

	err = dl.Data.DeathRecords().CreateMultiple(records)
	if err != nil{
		return err
	}

	return nil
}



// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (dl *DeathLoader) downloadDeathFiles(filepath string) ([]string, error) {

	client := &http.Client{}
	// Download 2020 data
	reqBody := `{"Rows":[{"UniqueName":"[Death County]","Domain":"Demographics","ReportFieldSize":1,"OutputOrder":0,"VariableOfInterest":false,"LoadToDesigner":true,"Location":"ROW","SubReportHeader":"County of Residence","SubReportHeaderHeight":0.1,"SubReportFooter":"","SubReportFooterHeight":1,"IncludeSubtotals":false,"FilteredMembers":[]}],"Columns":[{"UniqueName":"[Death Month]","Domain":"Time/Location","ReportFieldSize":1,"OutputOrder":0,"VariableOfInterest":false,"LoadToDesigner":false,"Location":"ROW","SubReportHeader":"Month of Death","SubReportHeaderHeight":0.1,"SubReportFooter":"","SubReportFooterHeight":1,"IncludeSubtotals":false,"FilteredMembers":[]}],"Slicers":[{"CubeDimensionID":18053,"UniqueName":"[Death Year]","Promoted":false,"FilteredMembers":["2020 **"]}],"DatasetCode":"Mortality","ReportID":1813,"IncludeRowTotals":true,"IncludeColumnTotals":true,"IncludeGrandTotal":false,"IncludeTotalsHeader":false,"PercentCalcMethod":"rowtotal","PercentDimension":"SELECT","Measures":[{"Caption":"Deaths","Description":"","Name":"Deaths","UniqueName":"[Measures].[Deaths]","BuilderName":"Death Count","CubeName":"[Measures].[Deaths]","CubeMeasureID":0,"BuilderToolTip":"Death Count","IsCount":true,"IncludeValue":true,"IncludeRunningCount":false,"IncludePercent":false,"IncludeRunningPercent":false,"PercentDimension":"rowtotal","OutputFormat":"N0","QueryAttributes":null,"OutputColumns":null,"LoadToDesigner":true,"DesignerOrder":"1","Location":null,"ExcludeCountInTotalGroup":false,"ExcludeRunningCountInTotalGroup":false,"ExcludePercentInTotalGroup":false,"ExcludeRunningPercentInTotalGroup":false,"MeasureHeaderText":null,"MeasureFooterText":null}],"IncludeTotalHeaderBlock":false,"SummarizeInSubreports":true}`
	resp, err := dl.getCookieAndRequestFiles(reqBody, client)
	if err != nil {
		return []string{}, err
	}
	deaths2020Files, err := dl.DownloadFiles(filepath, resp, "2020deaths")
	if err != nil {
		return []string{}, err
	}

	_, err = os.Open("data/imports/otherDeaths/DataFile1.csv")
	if err == nil || !os.IsNotExist(err) {
		return append(deaths2020Files, "./data/imports/otherDeaths/DataFile1.csv"), nil
	}
	reqBody = `{"Rows":[{"UniqueName":"[Death County]","Domain":"Demographics","ReportFieldSize":1,"OutputOrder":0,"VariableOfInterest":false,"LoadToDesigner":true,"Location":"ROW","SubReportHeader":"County of Residence","SubReportHeaderHeight":0.1,"SubReportFooter":"","SubReportFooterHeight":1,"IncludeSubtotals":false,"FilteredMembers":[]}],"Columns":[{"UniqueName":"[Death Month]","Domain":"Time/Location","ReportFieldSize":1,"OutputOrder":0,"VariableOfInterest":false,"LoadToDesigner":false,"Location":"ROW","SubReportHeader":"Month of Death","SubReportHeaderHeight":0.1,"SubReportFooter":"","SubReportFooterHeight":1,"IncludeSubtotals":false,"FilteredMembers":[]}],"Slicers":[{"CubeDimensionID":18053,"UniqueName":"[Death Year]","Promoted":false,"FilteredMembers":["2015","2016","2017","2018","2019 **"]}],"DatasetCode":"Mortality","ReportID":1813,"IncludeRowTotals":true,"IncludeColumnTotals":true,"IncludeGrandTotal":false,"IncludeTotalsHeader":false,"PercentCalcMethod":"rowtotal","PercentDimension":"SELECT","Measures":[{"Caption":"Deaths","Description":"","Name":"Deaths","UniqueName":"[Measures].[Deaths]","BuilderName":"Death Count","CubeName":"[Measures].[Deaths]","CubeMeasureID":0,"BuilderToolTip":"Death Count","IsCount":true,"IncludeValue":true,"IncludeRunningCount":false,"IncludePercent":false,"IncludeRunningPercent":false,"PercentDimension":"rowtotal","OutputFormat":"N0","QueryAttributes":null,"OutputColumns":null,"LoadToDesigner":true,"DesignerOrder":"1","Location":null,"ExcludeCountInTotalGroup":false,"ExcludeRunningCountInTotalGroup":false,"ExcludePercentInTotalGroup":false,"ExcludeRunningPercentInTotalGroup":false,"MeasureHeaderText":null,"MeasureFooterText":null}],"IncludeTotalHeaderBlock":false,"SummarizeInSubreports":true}`
	resp, err = dl.getCookieAndRequestFiles(reqBody, client)
	if err != nil {
		return []string{}, err
	}
	otherDeathFiles, err := dl.DownloadFiles(filepath, resp, "otherDeaths")
	if err != nil {
		return []string{}, err
	}

	return append(deaths2020Files, otherDeathFiles...), nil
}

func (dl *DeathLoader) getCookieAndRequestFiles(reqBody string, client *http.Client) (*http.Response, error) {
	cookieReq, _ := http.NewRequest("POST", "http://publicapps.odh.ohio.gov/EDW/Reports/RequestReportData", strings.NewReader(reqBody))
	cookieReq.Header.Set("Content-Type", "application/json; charset=UTF-8")
	cookieResp, err := client.Do(cookieReq)
	if err != nil {
		return nil, err
	}

	sessionCookie := cookieResp.Header.Get("Session")
	sessionCookie = strings.Trim(sessionCookie, "[]")
	sessionCookieParts := strings.Split(sessionCookie, ":")
	fmt.Println(sessionCookieParts)
	req, _ := http.NewRequest("GET", "http://publicapps.odh.ohio.gov/EDW/Reports/GetDelimitedReportData", nil)
	req.Header.Add("Cookie", "ASP.NET_SessionId="+sessionCookieParts[2]+";")
	resp, err := client.Do(req)
	return resp, err
}

func (dl *DeathLoader) DownloadFiles(filepath string, resp *http.Response, extractFolder string) ([]string, error) {
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return []string{}, err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return Unzip(filepath, "./data/imports/"+extractFolder+"/")
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
	return filenames[:len(filenames)-1], nil
}
