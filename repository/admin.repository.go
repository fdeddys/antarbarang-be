package repository

import (
	"errors"

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
		SELECT id, nama, status, last_update_by, last_update
		FROM public.admins
		WHERE id = $1;
	`
	var admin model.Admin
	err := db().
		QueryRow(sqlStatement, id).
		Scan(
			&admin.ID,
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

func SaveAdmin(admin model.Admin) (int64, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()
	lastInsertId := int64(0)

	kode, errKode := generateKode(constanta.PREFIX_ADMIN)
	admin.Kode = kode
	admin.Status = enumerate.ACTIVE
	if errKode != nil {
		return lastInsertId, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO public.admins
			(nama, kode, status, last_update_by, last_update)
		VALUES ($1::text, $2, $3, $4, $5)
		RETURNING id`

	err := db.QueryRow(
		sqlStatement,
		admin.Nama, admin.Kode, admin.Status, dto.CurrUser, currTime).
		Scan(&lastInsertId)
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil
}

func FindAAdminByCode(kode string) (model.Admin, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, kode, nama, status, last_update_by, last_update
		FROM public.admins;		
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
