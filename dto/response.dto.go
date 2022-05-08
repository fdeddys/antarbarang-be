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
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type SearchResultDto struct {
	TotalRow int         `json:"totalRow"`
	Page     int         `json:"page"`
	Count    int         `json:"count"`
	Contents interface{} `json:"contents"`
	Error    string      `json:"error"`
}
