package dto

var CurrUser string
var CurrUserId int64

type ContentResponse struct {
	ErrCode  string      `json:"errCode"`
	ErrDesc  string      `json:"errDesc"`
	Contents interface{} `json:"contents"`
}

type LoginRequestDto struct {
	Kode     string `json:"kode"`
	Password string `json:"password"`
}

type LoginUsernameDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	ErrCode string `json:"errCode"`
	ErrDesc string `json:"errDesc"`
	Token   string `json:"token"`
}

type SearchRequestDto struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	SellerId string `json:"sellerId"`
}

type SearchResultDto struct {
	TotalRow int         `json:"totalRow"`
	Page     int         `json:"page"`
	Count    int         `json:"count"`
	Contents interface{} `json:"contents"`
	Error    string      `json:"error"`
}

type SearchTransaksiRequestDto struct {
	SellerName   string `json:"sellerName"`
	DriverName   string `json:"driverName"`
	CustomerName string `json:"customerName"`
	Status       string `json:"status"`
	Tgl1         string `json:"tgl1"`
	Tgl2         string `json:"tgl2"`
	DriverID     string `json:"driverId"`
}

type ChangePasswordRequestDto struct {
	DriverId    int    `json:"driverId"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
}
