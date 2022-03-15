package services

import (
	"context"
	"fmt"
	"time"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

type DriverService struct {
}

func (d *DriverService) CreateDriver(driver model.Driver) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	currTime := time.Now().Unix()
	db := database.GetConn

	sqlStatement := `
		INSERT INTO public.drivers
		(nama, hp, alamat, photo, status, last_update_by, last_update)
		VALUES ($1::text, $2::text, $3::text, $4::text, 0, $5, $6::bigint)
		RETURNING id`

	fmt.Println("name", driver.Name, "addr", driver.Address, "pict", driver.Picture, "status", driver.Status, "curtime", currTime, "cur user", dto.CurrUser)

	lastInsertId := 0
	err := db().
		QueryRow(context.Background(), sqlStatement, driver.Name, driver.Address, driver.Picture, driver.Status, dto.CurrUser, currTime).
		Scan(&lastInsertId)

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

	driver, err := repository.FindById(id)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = driver
	return result
}
