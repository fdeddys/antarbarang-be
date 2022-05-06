package services

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

// ParameterService ...
type ParameterService struct {
}

// GetDataOrderById ...
func (p ParameterService) GetByName(paramName string) model.Parameter {

	// var res dbmodels.Parameter
	// var err error
	res, errCode, _, _ := repository.GetParameterByNama(paramName)
	if errCode == constanta.ERR_CODE_00 {
		return res
	}

	return model.Parameter{ID: 0, Nama: "", Value: ""}
}
