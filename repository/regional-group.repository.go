package repository

import (
	"context"
	"fmt"
	"time"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindAllRegionalGroup() ([]model.RegionalGroup, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, status, last_update_by, last_update
		FROM regional_group;
	`
	var regionalGroups []model.RegionalGroup
	datas, err := db().
		QueryContext(
			context.Background(),
			sqlStatement)

	for datas.Next() {
		var regionalGroup model.RegionalGroup

		datas.Scan(
			&regionalGroup.ID,
			&regionalGroup.Nama,
			&regionalGroup.Status,
			&regionalGroup.LastUpdateBy,
			&regionalGroup.LastUpdate,
		)
		regionalGroups = append(regionalGroups, regionalGroup)

	}
	if err != nil {
		return regionalGroups, err
	}
	return regionalGroups, nil

}

func SaveRegionalGroup(regionalGroup model.RegionalGroup) (model.RegionalGroup, error) {

	// currTime := util.GetCurrDate().Format("2006-01-02 15:04:05")
	currTime := util.GetCurrTimeString()
	db := database.GetConn()

	regionalGroup.Status = enumerate.ACTIVE

	sqlStatement := `
		INSERT INTO regional_group
			(nama, status, last_update_by, last_update)
		VALUES (?, ?, ?, ?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return regionalGroup, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		regionalGroup.Nama, regionalGroup.Status, dto.CurrUser, currTime)
	if err != nil {
		return regionalGroup, err
	}
	idGenerated, err := res.LastInsertId()
	regionalGroup.ID = idGenerated
	return regionalGroup, nil
}

func UpdateRegionalGroup(regionalGroup model.RegionalGroup) (string, error) {

	currTime := util.GetCurrTimeString()
	db := database.GetConn()

	sqlStatement := `
		UPDATE regional_group
		SET nama=?, status=?, last_update_by=?, last_update=?
		WHERE id=?;
	`

	res, err := db.Exec(
		sqlStatement,
		regionalGroup.Nama, regionalGroup.Status, dto.CurrUser, currTime, regionalGroup.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), nil
}
