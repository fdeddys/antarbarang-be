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

func FindAdminById(id int) (model.Admin, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, kode, status, last_update_by, last_update
		FROM public.admins
		WHERE id = $1;
	`
	var admin model.Admin
	err := db().
		QueryRow(sqlStatement, id).
		Scan(
			&admin.ID,
			&admin.Nama,
			&admin.Kode,
			&admin.Status,
			&admin.LastUpdateBy,
			&admin.LastUpdate,
		)
	admin.LastUpdateStr = util.DateUnixToString(admin.LastUpdate)
	if err != nil {
		return admin, err
	}
	return admin, nil

}

func SaveAdmin(admin model.Admin) (model.Admin, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()
	lastInsertId := int64(0)

	fmt.Println("Start generate kode")
	kode, errKode := generateKode(constanta.PREFIX_ADMIN)
	admin.Kode = kode
	admin.Status = enumerate.ACTIVE
	if errKode != nil {
		return admin, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}
	admin.ID = lastInsertId

	sqlStatement := `
		INSERT INTO public.admins
			(nama, Password, kode, status, last_update_by, last_update)
		VALUES ($1::text, $2, $3, $4, $5, $6)
		RETURNING id`

	fmt.Println("Start query row")
	err := db.QueryRow(
		sqlStatement,
		admin.Nama, admin.Password, admin.Kode, admin.Status, dto.CurrUser, currTime).
		Scan(&lastInsertId)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func FindAdminByCode(kode string) (model.Admin, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, kode, nama, status, last_update_by, last_update
		FROM public.admins	
		WHERE kode = $1;
	`
	var admin model.Admin
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&admin.ID,
			&admin.Kode,
			&admin.Nama,
			&admin.Status,
			&admin.LastUpdateBy,
			&admin.LastUpdate,
		)
	admin.LastUpdateStr = util.DateUnixToString(admin.LastUpdate)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func UpdateAdmin(admin model.Admin) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.admins
		SET nama=$1, last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		admin.Nama, dto.CurrUser, currTime, admin.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), nil
}

func LoginAdminByCode(kode string) (model.Admin, error) {
	db := database.GetConn()

	sqlStatement := `
		SELECT nama, password, status
		FROM public.admins
		WHERE kode = $1; 
	`
	var admin model.Admin
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&admin.Nama,
			&admin.Password,
			&admin.Status,
		)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
