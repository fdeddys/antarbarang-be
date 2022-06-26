package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func generateQueryRegional(searchRequestDto dto.SearchRequestDto, page, limit int) (string, string) {

	var kriteriaRegionalName = "%"
	if searchRequestDto.Nama != "" {
		kriteriaRegionalName += searchRequestDto.Nama + "%"
	}

	sqlFind := fmt.Sprintf(`
		SELECT r.id, r.nama, r.status, r.last_update_by, r.last_update, regional_group_id, rg.nama 
		FROM regional r
		left join regional_group rg on r.regional_group_id  = rg.id 		
		WHERE r.nama like '%v'  
		ORDER BY rg.nama, r.nama

		LIMIT %v, %v
		`,
		kriteriaRegionalName,
		((page - 1) * limit), limit)

	sqlCount := `
		SELECT count(*)
		FROM regional	
		WHERE nama like '%'  
	`
	return sqlFind, sqlCount

}

func GetRegionalPage(searchRequestDto dto.SearchRequestDto, page, count int) ([]model.Regional, int, error) {
	db := database.GetConn()
	var regionals []model.Regional
	var total int

	sqlFind, sqlCount := generateQueryRegional(searchRequestDto, page, count)

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerySearchRegional(db, sqlFind, &regionals, errQuery)
	go AsyncQueryCount(db, sqlCount, &total, errCount)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return regionals, 0, resErrCount
	}

	if resErrQuery != nil {
		return regionals, 0, resErrQuery
	}

	return regionals, total, nil
}

func AsyncQuerySearchRegional(db *sql.DB, sqlFind string, regionals *[]model.Regional, resChan chan error) {

	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		resChan <- err
		return
	}
	// r.id, r.nama, r.status, r.last_update_by, r.last_update, regional_group_id, rg.nama
	for datas.Next() {
		var regional model.Regional
		err = datas.Scan(
			&regional.ID,
			&regional.Nama,
			&regional.Status,
			&regional.LastUpdateBy,
			&regional.LastUpdate,
			&regional.RegionalGroupId,
			&regional.RegionalGroupName,
		)
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*regionals = append(*regionals, regional)
	}
	resChan <- nil

}

func FindRegionalByRegionalGroupID(regionalGroupId int) (model.Regional, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, status, last_update_by, last_update
		FROM regional where regional_group_id = ?;
	`
	var regional model.Regional
	err := db().
		QueryRow(sqlStatement, regionalGroupId).
		Scan(
			&regional.ID,
			&regional.Nama,
			&regional.Status,
			&regional.LastUpdateBy,
			&regional.LastUpdate,
			&regional.RegionalGroupId,
		)
	if err != nil {
		return regional, err
	}
	return regional, nil

}

func SaveRegional(regional model.Regional) (model.Regional, error) {

	currTime := util.GetCurrTimeString()
	db := database.GetConn()

	regional.Status = enumerate.ACTIVE

	sqlStatement := `
	INSERT INTO regional
		(nama, status, last_update_by, last_update, regional_group_id)
		VALUES (?, ?, ?, ?, ?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return regional, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		regional.Nama, regional.Status, dto.CurrUser, currTime, regional.RegionalGroupId)
	if err != nil {
		return regional, err
	}
	idGenerated, err := res.LastInsertId()
	regional.ID = idGenerated
	return regional, nil
}

func UpdateRegional(regional model.Regional) (string, error) {

	currTime := util.GetCurrTimeString()
	db := database.GetConn()

	sqlStatement := `
		UPDATE regional
		SET nama=?, status=?, last_update_by=?, last_update=?, regional_group_id = ?
		WHERE id=?;
	`

	res, err := db.Exec(
		sqlStatement,
		regional.Nama, regional.Status, dto.CurrUser, currTime, regional.RegionalGroupId, regional.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), nil
}
