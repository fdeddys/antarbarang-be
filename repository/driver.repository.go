package repository

import (
	"errors"
	"fmt"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
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
		Scan(
			&driver.ID,
			&driver.Nama,
			&driver.Hp,
			&driver.Alamat,
			&driver.Photo,
			&driver.Status,
			&driver.LastUpdateBy,
			&driver.LastUpdate,
		)
	driver.LastUpdateStr = util.DateUnixToString(driver.LastUpdate)
	if err != nil {
		return driver, err
	}
	return driver, nil

}

func SaveDriver(driver model.Driver) (int64, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()
	lastInsertId := int64(0)

	kode, errKode := generateKode(constanta.PREFIX_DRIVER)
	driver.Kode = kode
	driver.Status = enumerate.ACTIVE
	if errKode != nil {
		return lastInsertId, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO public.drivers
			(nama, kode, hp, alamat, photo, status, last_update_by, last_update)
		VALUES ($1::text, $2::text, $3::text, $4::text, $5, $6, $7, $8)
		RETURNING id`

	err := db.QueryRow(
		sqlStatement,
		driver.Nama, driver.Kode, driver.Hp, driver.Alamat, driver.Photo, driver.Status, dto.CurrUser, currTime).
		Scan(&lastInsertId)
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}

func FindDriverByCode(kode string) (model.Driver, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, nama, kode, hp, alamat, photo, status, last_update_by, last_update
		FROM public.drivers	
		WHERE kode = $1;
	`
	var driver model.Driver
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&driver.ID,
			&driver.Nama,
			&driver.Kode,
			&driver.Hp,
			&driver.Alamat,
			&driver.Photo,
			&driver.Status,
			&driver.LastUpdateBy,
			&driver.LastUpdate,
		)
	driver.LastUpdateStr = util.DateUnixToString(driver.LastUpdate)
	if err != nil {
		return driver, err
	}
	return driver, nil
}

func LoginDriverByCode(kode string) (model.Driver, error) {
	db := database.GetConn()

	sqlStatement := `
		SELECT nama, password, status
		FROM public.drivers
		WHERE kode = $1; 
	`
	var driver model.Driver
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&driver.Nama,
			&driver.Password,
			&driver.Status,
		)
	if err != nil {
		return driver, err
	}
	return driver, nil
}

func UpdateDriver(driver model.Driver) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.drivers
		SET nama=$1,  last_update_by=$2, last_update=$3, hp=$4, alamat=$5
		WHERE id=$6;
	`

	res, err := db.Exec(
		sqlStatement,
		driver.Nama, dto.CurrUser, currTime, driver.Hp, driver.Alamat, driver.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func UpdateStatusDriver(idDriver int64, statusDriver interface{}) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.drivers
		SET status=$1,  last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		statusDriver, dto.CurrUser, currTime, idDriver)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func ChangePasswordDriver(driver model.Driver) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.drivers
		SET password=$1,  last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		driver.Password, dto.CurrUser, currTime, driver.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}
