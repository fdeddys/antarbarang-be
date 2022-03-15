package repository

import (
	"context"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/model"
)

func FindById(id int) (model.Driver, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, hp, alamat, photo, status, last_update_by, last_update
		FROM public.drivers
		WHERE id = $1;
	`
	var driver model.Driver
	err := db().
		QueryRow(context.Background(), sqlStatement, id).
		Scan(&driver)
	if err != nil {
		return driver, err
	}
	return driver, nil

}
