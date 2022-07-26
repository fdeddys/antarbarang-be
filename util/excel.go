package util

import (
	"fmt"
	"time"

	"com.ddabadi.antarbarang/dto"

	"github.com/xuri/excelize/v2"
)

func ExportToExcelReportTransaksi(reportPayements []dto.ReportTransaksi, dateStart, dateEnd, namaFile string) (filename string, success bool) {

	filename = fmt.Sprintf("%v_%v_%v.xlsx", namaFile, dateStart, dateEnd)

	sheet1Name := "Sheet1"
	xls := excelize.NewFile()
	index := xls.NewSheet(sheet1Name)

	t1, _ := time.Parse("2006-01-02", dateStart)
	t2, _ := time.Parse("2006-01-02", dateEnd)
	fmt.Println("Waktu nya adalah", t1.Format("02-January-2006"))
	no := 1
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), "REPORT TRANSAKSI")

	no++
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), fmt.Sprintf("%v - %v ", t1.Format("02-January-2006"), t2.Format("02-January-2006")))

	no = no + 2
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), "#")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), "TanggalAntar")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), "Keterangan")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), "SellerRegion")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), "SellerName")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), "CustomerRegion")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), "CustomerName")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), "DriverName")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), "Total")
	urut := 0

	for _, rs := range reportPayements {

		// fmt.Println(rs.ReceiveNo, "-", rs.ReceiveTgl)
		no++
		urut++
		xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), urut)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), rs.TanggalAntar)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), rs.Keterangan)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), rs.SellerRegion)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), rs.SellerName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), rs.CustomerRegion)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), rs.CustomerName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), rs.DriverName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), rs.Biaya)
	}

	xls.SetActiveSheet(index)
	if err := xls.SaveAs(filename); err != nil {
		fmt.Println("Erorr create = ", err.Error())
		success = false
		return
	}
	success = true
	return
}
