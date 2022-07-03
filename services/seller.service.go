package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type SellerService struct {
}

func (d *SellerService) CreateSeller(seller model.Seller) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveSeller(seller)
	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (d *SellerService) GetSellerByID(id int64) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.FindSellerById(id)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = seller
	return result
}

func (d *SellerService) GetSellerByKode(kode string) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.FindSellerByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = seller
	return result
}

func (s *SellerService) LoginSellerByKode(kode, password string) dto.LoginResponseDto {
	var result dto.LoginResponseDto
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.LoginSellerByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Token = ""
		return result
	}

	if seller.Status != enumerate.ACTIVE {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Token = ""
		return result
	}

	if seller.Password != password {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Token = ""
		return result
	}

	result.Token = generateToken(seller.Nama, seller.ID)
	return result
}

func (s *SellerService) UpdateSeller(seller model.Seller) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateSeller(seller)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (s *SellerService) UpdateStatusSellerActive(sellerId int64, active bool) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	statusSeller := enumerate.NONACTIVE
	if active {
		statusSeller = enumerate.ACTIVE
	}

	msg, err := repository.UpdateStatusSeller(sellerId, statusSeller)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (s *SellerService) ChangePasswordSeller(seller model.Seller) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.ChangePasswordSeller(seller)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (s *SellerService) SearchSellerPage(searchRequestDto dto.SearchRequestDto, page, count int) dto.SearchResultDto {

	var result dto.SearchResultDto
	data, totalData, err := repository.GetSellerPage(searchRequestDto, page, count)

	if err != nil {
		result.Error = err.Error()
		return result
	}
	result.Contents = data
	result.TotalRow = totalData
	result.Page = page
	result.Count = len(data)

	return result
}
