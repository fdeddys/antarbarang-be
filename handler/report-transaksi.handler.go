package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/services"
)

var reportService = new(services.ReportTransaksiService)

func GetReportTransaksi(w http.ResponseWriter, r *http.Request) {

	req := dto.FilterReportDate{}
	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &req)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}

	filename, success := reportService.GenerateReportTransaksi(req)
	if success {
		fmt.Println("download EXCEL ")
		w.Header().Set("content-type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.WriteHeader(http.StatusOK)

		// file, _ := os.Open(filename)
		fileRpt, _ := os.ReadFile(filename)
		w.Write(fileRpt)

		os.Remove(filename)
	}

	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
}
