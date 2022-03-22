package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type DriverService struct {
}

func (d *DriverService) CreateDriver(driver model.Driver) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveDriver(driver)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (d *DriverService) GetDriverByID(id int) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	driver, err := repository.FindDriverById(id)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = driver
	return result
}

func (d *DriverService) GetDriverByKode(kode string) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.FindDriverByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = seller
	return result
}

func (d *DriverService) LoginDriverByKode(kode, password string) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.LoginDriverByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}

	if seller.Status != enumerate.ACTIVE {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Contents = "User status not active !"
		return result
	}

	if seller.Password != password {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Contents = "Password not match !"
		return result
	}

	result.Contents = "Login success"
	return result
}
