package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type RegionalGroupService struct {
}

func (d *RegionalGroupService) CreateRegionalGroup(regionalGroup model.RegionalGroup) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveRegionalGroup(regionalGroup)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (d *RegionalGroupService) UpdateRegionalGroup(regionalGroup model.RegionalGroup) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateRegionalGroup(regionalGroup)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (d *RegionalGroupService) GetAllRegionalGroup() dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	regionals, err := repository.FindAllRegionalGroup()
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = regionals
	return result
}
