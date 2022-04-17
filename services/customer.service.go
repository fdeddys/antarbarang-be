package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type CustomerService struct {
}

func (c *CustomerService) CreateCustomer(customer model.Customer) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveCustomer(customer)
	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (c *CustomerService) GetCustomerByID(custId int64) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	customer, err := repository.FindCustomerById(int(custId))
	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		return result
	}
	result.Contents = customer
	return result

}

func (c *CustomerService) GetCustomerByNama(nama string) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	customers, err := repository.FindCustomerByNama(nama)
	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		return result
	}
	result.Contents = customers
	return result

}

func (c *CustomerService) GetCustomerBySellerId(sellerId int64) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	customers, err := repository.FindCustomerBySellerId(sellerId)
	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		return result
	}
	result.Contents = customers
	return result

}

func (c *CustomerService) UpdateCustomer(customer model.Customer) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateCustomer(customer)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_12
		result.ErrDesc = constanta.ERR_CODE_12_FAILED_UPDATE_DATA
		return result
	}
	result.Contents = msg
	return result
}

func (c *CustomerService) UpdateStatusCustomerActive(customerId int64, active bool) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	statusDriver := enumerate.NONACTIVE
	if active {
		statusDriver = enumerate.ACTIVE
	}

	msg, err := repository.UpdateStatusCustomer(customerId, statusDriver)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_12
		result.ErrDesc = constanta.ERR_CODE_12_FAILED_UPDATE_DATA
		return result
	}
	result.Contents = msg
	return result
}
