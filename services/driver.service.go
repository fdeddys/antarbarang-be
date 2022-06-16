package services

import (
	"fmt"

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

func (d *DriverService) LoginDriverByKode(kode, password string) dto.LoginResponseDto {
	var result dto.LoginResponseDto
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	driver, err := repository.LoginDriverByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		// result.Contents = err.Error()
		result.Token = ""
		return result
	}

	if driver.Status != enumerate.ACTIVE {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		// result.Contents = "User status not active !"
		result.Token = ""
		return result
	}

	if driver.Password != password {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		// result.Contents = "Password not match !"
		result.Token = ""
		return result
	}

	result.Token = generateToken(driver.Nama, driver.ID)
	return result
}

func (d *DriverService) UpdateDriver(driver model.Driver) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateDriver(driver)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (d *DriverService) UpdateStatusDriverActive(driverId int64, active bool) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	statusDriver := enumerate.NONACTIVE
	if active {
		statusDriver = enumerate.ACTIVE
	}

	msg, err := repository.UpdateStatusDriver(driverId, statusDriver)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (d *DriverService) ChangePasswordDriver(changeReqModel dto.ChangePasswordRequestDto) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	password, errDr := repository.FindPasswordDriverById(changeReqModel.DriverId)
	if errDr != nil {
		result.Contents = errDr.Error()
		result.ErrCode = constanta.ERR_CODE_12
		result.ErrDesc = constanta.ERR_CODE_12_FAILED_UPDATE_DATA
		return result
	}

	fmt.Println("old :", changeReqModel.OldPassword, ":new pass:", password)
	if changeReqModel.OldPassword != password {
		result.Contents = "old password not match !"
		result.ErrCode = constanta.ERR_CODE_12
		result.ErrDesc = constanta.ERR_CODE_12_FAILED_UPDATE_DATA
		return result
	}

	msg, err := repository.ChangePasswordDriver(changeReqModel)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_12
		result.ErrDesc = constanta.ERR_CODE_12_FAILED_UPDATE_DATA
		return result
	}
	result.Contents = msg
	return result
}

func (d *DriverService) SearchDriverPage(searchRequestDto dto.SearchRequestDto, page, count int) dto.SearchResultDto {

	var result dto.SearchResultDto
	data, totalData, err := repository.GetDriverPage(searchRequestDto, page, count)

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
