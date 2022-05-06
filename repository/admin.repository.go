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

func FindAdminById(id int) (model.Admin, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, kode, status, last_update_by, last_update
		FROM admins
		WHERE id = ?;
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

	fmt.Println("Start generate kode")
	kode, errKode := generateKode(constanta.PREFIX_ADMIN)
	fmt.Println("Kode  => ", kode)
	admin.Kode = kode
	admin.Status = enumerate.ACTIVE
	if errKode != nil {
		return admin, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO admins
			(nama, Password, kode, status, last_update_by, last_update)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return admin, err
	}
	defer stmt.Close()

	fmt.Println("Kode genetated => ", admin.Kode)
	res, err := stmt.ExecContext(
		ctx,
		admin.Nama, admin.Password, admin.Kode, admin.Status, dto.CurrUser, currTime)
	if err != nil {
		return admin, err
	}
	idGenerated, err := res.LastInsertId()
	admin.ID = idGenerated
	return admin, nil
}

func FindAdminByCode(kode string) (model.Admin, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, kode, nama, status, last_update_by, last_update
		FROM admins	
		WHERE kode = ?;
	`
	var admin model.Admin
	err := db.QueryRow(
		sqlStatement, kode).
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
		UPDATE admins
		SET nama=?, last_update_by=?, last_update=?
		WHERE id=?;
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

func LoginAdminByNama(nama string) (model.Admin, error) {
	db := database.GetConn()

	sqlStatement := `
		SELECT id, nama, password, status
		FROM admins
		WHERE nama = ? and status =1 ; 
	`
	var admin model.Admin
	err := db.
		QueryRow(sqlStatement, nama).
		Scan(
			&admin.ID,
			&admin.Nama,
			&admin.Password,
			&admin.Status,
		)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func ChangePasswordAdmin(admin model.Admin) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE public.admin
		SET password=$1,  last_update_by=$2, last_update=$3
		WHERE id=$4;
	`

	res, err := db.Exec(
		sqlStatement,
		admin.Password, dto.CurrUser, currTime, admin.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}
