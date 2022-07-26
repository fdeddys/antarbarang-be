package dto

type UpdateLngLatCustomerRequestDto struct {
	CustId int64  `json:"custId"`
	Lng    string `json:"lng"`
	Lat    string `json:"lat"`
}
