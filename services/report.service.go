package services

import (
	"fmt"

	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/repository"
	"com.ddabadi.antarbarang/util"
)

type ReportTransaksiService struct {
}

// Approve ...
func (o ReportTransaksiService) GenerateReportTransaksi(filterData dto.FilterReportDate) (filename string, success bool) {

	dateStart := filterData.StartDate + " 00:00:00"
	dateEnd := filterData.EndDate + " 23:59:59"
	datas := generateDataReport(dateStart, dateEnd)
	// filename = ExportToCSV(datas, filterData.StartDate, filterData.EndDate, "report-payment")

	filename, success = util.ExportToExcelReportTransaksi(datas, filterData.StartDate, filterData.EndDate, "report-transaksi")
	fmt.Println("success =>", success)
	return
}

func generateDataReport(dateStart, dateEnd string) []dto.ReportTransaksi {

	datas := repository.ReportTransaksiByDate(dateStart, dateEnd)

	return datas
}
