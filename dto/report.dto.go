package dto

type FilterReportDate struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type ReportTransaksi struct {
	TanggalAntar   string
	Keterangan     string
	Biaya          float32
	SellerRegion   string
	SellerId       int64
	SellerName     string
	CustomerRegion string
	CustomerId     int64
	CustomerName   string
	DriverId       int64
	DriverName     string
}
