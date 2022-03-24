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

func FindSellerById(id int64) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, nama, hp, alamat, status, last_update_by, last_update
		FROM public.sellers
		WHERE id = $1;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, id).
		Scan(
			&seller.ID,
			&seller.Nama,
			&seller.Hp,
			&seller.Alamat,
			&seller.Status,
			&seller.LastUpdateBy,
			&seller.LastUpdate,
		)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func FindSellerByCode(kode string) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, nama, hp, alamat, status, last_update_by, last_update
		FROM public.sellers
		WHERE kode = $1;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&seller.ID,
			&seller.Nama,
			&seller.Hp,
			&seller.Alamat,
			&seller.Status,
			&seller.LastUpdateBy,
			&seller.LastUpdate,
		)
	seller.LastUpdateStr = util.DateUnixToString(seller.LastUpdate)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func SaveSeller(seller model.Seller) (model.Seller, error) {

	lastInsertId := 0
	db := database.GetConn
	kode, errKode := generateKode(constanta.PREFIX_SELLER)
	seller.Kode = kode
	seller.Status = enumerate.ACTIVE
	if errKode != nil {
		return model.Seller{}, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO public.sellers
			(nama, hp, kode, password, alamat,  status, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db().
		QueryRow(
			sqlStatement,
			seller.Nama, seller.Hp, seller.Kode, seller.Password, seller.Alamat, seller.Status, dto.CurrUser, util.GetCurrTimeUnix(),
		).
		Scan(&lastInsertId)

	if err != nil {
		return seller, err
	}
	seller.ID = int64(lastInsertId)
	return seller, nil
}

func LoginSellerByCode(kode string) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT nama, password, status
		FROM public.sellers
		WHERE kode = $1; 
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&seller.Nama,
			&seller.Password,
			&seller.Status,
		)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func UpdateSeller(seller model.Seller) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.sellers
		SET nama=$1,  last_update_by=$2, last_update=$3, hp=$4, alamat=$5
		WHERE id=$6;
	`

	res, err := db.Exec(
		sqlStatement,
		seller.Nama, dto.CurrUser, currTime, seller.Hp, seller.Alamat, seller.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func UpdateStatusSeller(idSeller int64, statusSeller interface{}) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.sellers
		SET status=$1,  last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		statusSeller, dto.CurrUser, currTime, idSeller)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func ChangePasswordSeller(seller model.Seller) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.sellers
		SET password=$1,  last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		seller.Password, dto.CurrUser, currTime, seller.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}
