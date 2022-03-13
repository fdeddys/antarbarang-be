package dto

var CurrUser string

type ContentResponse struct {
	ErrCode  string      `json:"errCode"`
	ErrDesc  string      `json:"errDesc"`
	Contents interface{} `json:"contents"`
}
