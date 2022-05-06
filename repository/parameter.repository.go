package repository

import (
	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/model"
)

func GetParameterByNama(nama string) (model.Parameter, string, string, error) {

	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, value, isviewable, last_update_by, last_update
		FROM parameter
		WHERE nama = ?;
	`
	var parameter model.Parameter
	err := db().
		QueryRow(sqlStatement, nama).
		Scan(
			&parameter.ID,
			&parameter.Nama,
			&parameter.Value,
			&parameter.IsViewable,
			&parameter.LastUpdateBy,
			&parameter.LastUpdate,
		)
	if err != nil {
		return parameter, constanta.ERR_CODE_11, constanta.ERR_CODE_11_FAILED_GET_DATA, err
	}
	return parameter, constanta.ERR_CODE_00, constanta.ERR_CODE_00_MSG, nil

}
