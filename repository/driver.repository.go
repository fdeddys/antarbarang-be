package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

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
		FROM drivers
		WHERE id = ?;
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
		INSERT INTO drivers
			(nama, kode, hp, alamat, photo, password, status, last_update_by, last_update)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return 0, err
	}

	resp, err := stmt.ExecContext(
		ctx,
		driver.Nama,
		driver.Kode,
		driver.Hp,
		driver.Alamat,
		driver.Photo,
		driver.Password,
		driver.Status,
		dto.CurrUser,
		currTime,
	)
	if err != nil {
		return 0, err
	}

	lastInsertId, err = resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func FindDriverByCode(kode string) (model.Driver, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, nama, kode, hp, alamat, photo, status, last_update_by, last_update
		FROM drivers	
		WHERE kode = ?;
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
		FROM drivers
		WHERE kode = ?; 
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
		UPDATE drivers
		SET nama = ?,  last_update_by = ?, last_update = ?, hp = ?, alamat = ?
		WHERE id = ?;
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
		UPDATE drivers
		SET status = ?,  last_update_by= ?, last_update=?
		WHERE id   = ?;
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
		UPDATE drivers
		SET password= ?,  last_update_by= ?, last_update= ?
		WHERE id = ?;
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

func generateQueryDriver(searchRequestDto dto.SearchRequestDto, page, limit int) (string, string) {

	sqlFind := `
		SELECT id, nama, kode,  hp, alamat, photo, status, last_update_by, last_update
		FROM drivers		
		WHERE nama like '%' and kode like '%' 
	`
	sqlCount := `
		SELECT count(*)
		FROM drivers	
		WHERE nama like '%' and kode like '%' 
	`
	return sqlFind, sqlCount

}

func GetDriverPage(searchRequestDto dto.SearchRequestDto, page, count int) ([]model.Driver, int, error) {
	db := database.GetConn()
	var drivers []model.Driver
	var total int

	sqlFind, sqlCount := generateQueryDriver(searchRequestDto, page, count)

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerySearchDriver(db, sqlFind, &drivers, errQuery)
	go AsyncQueryCount(db, sqlCount, &total, errCount)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return drivers, 0, resErrCount
	}

	if resErrQuery != nil {
		return drivers, 0, resErrQuery
	}

	return drivers, total, nil
}

func AsyncQuerySearchDriver(db *sql.DB, sqlFind string, drivers *[]model.Driver, resChan chan error) {

	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		resChan <- err
		return
	}

	for datas.Next() {
		var driver model.Driver
		err = datas.Scan(
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
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*drivers = append(*drivers, driver)
	}
	resChan <- nil

}
