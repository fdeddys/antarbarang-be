package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type RegionalService struct {
}

func (d *RegionalService) CreateRegional(regional model.Regional) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveRegional(regional)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (d *RegionalService) UpdateRegional(regional model.Regional) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateRegional(regional)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (d *RegionalService) GetRegionalByRegionalGroupId(regionalGroupId int) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	driver, err := repository.FindRegionalByRegionalGroupID(regionalGroupId)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = driver
	return result
}

func (d *RegionalService) SearchRegionalPage(searchRequestDto dto.SearchRequestDto, page, count int) dto.SearchResultDto {

	var result dto.SearchResultDto
	data, totalData, err := repository.GetRegionalPage(searchRequestDto, page, count)

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
