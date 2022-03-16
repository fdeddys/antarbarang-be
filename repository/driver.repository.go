package repository

import (
	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindDriverById(id int) (model.Driver, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, hp, alamat, photo, status, last_update_by, last_update
		FROM public.drivers
		WHERE id = $1;
	`
	var driver model.Driver
	err := db().
		QueryRow(sqlStatement, id).
		Scan(&driver)
	if err != nil {
		return driver, err
	}
	return driver, nil

}

func SaveDriver(driver model.Driver) (int64, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()
	defer db.Close()

	sqlStatement := `
		INSERT INTO public.drivers
		(nama, hp, alamat, photo, status, last_update_by, last_update)
		VALUES ($1::text, $2::text, $3::text, $4::text, 0, $5, $6::bigint)
		RETURNING id`

	lastInsertId := 0
	err := db.QueryRow(
		sqlStatement, driver.Name, driver.Address, driver.Picture, driver.Status, dto.CurrUser, currTime).
		Scan(&lastInsertId)
	if err != nil {
		return int64(lastInsertId), err
	}
	return int64(lastInsertId), nil

}
