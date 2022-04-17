package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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
		SELECT id, kode, nama, hp, alamat, status, last_update_by, last_update
		FROM sellers
		WHERE id = ?;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, id).
		Scan(
			&seller.ID,
			&seller.Kode,
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
	seller.LastUpdateStr = util.DateUnixToString(seller.LastUpdate)
	return seller, nil
}

func FindSellerByCode(kode string) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, kode, nama, hp, alamat, status, last_update_by, last_update
		FROM sellers
		WHERE kode = ?;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&seller.ID,
			&seller.Kode,
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

	db := database.GetConn()
	kode, errKode := generateKode(constanta.PREFIX_SELLER)
	seller.Kode = kode
	seller.Status = enumerate.ACTIVE
	if errKode != nil {
		return model.Seller{}, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO sellers
			(nama, hp, kode, password, alamat,  status, last_update_by, last_update)
		VALUES (?,?,?,?,?,?,?,?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return model.Seller{}, err
	}

	res, err := stmt.ExecContext(ctx,
		seller.Nama, seller.Hp, seller.Kode, seller.Password, seller.Alamat, seller.Status, dto.CurrUser, util.GetCurrTimeUnix(),
	)

	if err != nil {
		return seller, err
	}
	seller.ID, _ = res.LastInsertId()
	return seller, nil
}

func LoginSellerByCode(kode string) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT nama, password, status
		FROM sellers
		WHERE kode = ?; 
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
		UPDATE sellers
		SET nama=?, last_update_by=?, last_update=?, hp=?, alamat=?
		WHERE id=?;
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
		UPDATE sellers
		SET status=?,  last_update_by=?, last_update=?
		WHERE id=?;
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
		UPDATE sellers
		SET password=?,  last_update_by=?, last_update=?
		WHERE id=?;
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
