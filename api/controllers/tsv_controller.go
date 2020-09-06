package controllers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Ilhamsat/go/api/responses"
)

type Urlapi struct {
	Hostname     string `json:"host"`
	Logname      string `json:"logname"`
	Time         string `json:"time"`
	Method       string `json:"method"`
	Urls         string `json:"url"`
	Responsecode string `json:"response"`
	Bytes        string `json:"bytes"`
	Referer      string `json:"referer"`
}

func (server *Server) Tsv(w http.ResponseWriter, r *http.Request) {

	csvFile, err := os.Open("./log_19950801.tsv")

	if err != nil {
		fmt.Println(err)
	}

	row1, err := bufio.NewReader(csvFile).ReadSlice('\n')
	if err != nil {
		fmt.Println(err)
	}
	_, err = csvFile.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t'

	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Urlapi
	var allRecords []Urlapi

	for _, each := range csvData {
		oneRecord.Hostname = each[0]
		oneRecord.Logname = each[1]
		oneRecord.Time = each[2]
		oneRecord.Method = each[3]
		oneRecord.Urls = each[4]
		oneRecord.Responsecode = each[5]
		oneRecord.Bytes = each[6]
		oneRecord.Referer = each[7]
		allRecords = append(allRecords, oneRecord)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responses.JSON(w, http.StatusOK, allRecords)
}
