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

func FindSellerById(id int64) (model.Seller, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, kode, nama, hp, alamat, status, regional_id, last_update_by, last_update
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
			&seller.RegionalId,
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
		SELECT id, kode, nama, hp, alamat, status, regional_id, last_update_by, last_update
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
			&seller.RegionalId,
			&seller.LastUpdateBy,
			&seller.LastUpdate,
		)
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
			(nama, hp, kode, password, alamat,  status, regional_id, last_update_by, last_update)
		VALUES (?,?,?,?,?,?,?,?,?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return model.Seller{}, err
	}

	res, err := stmt.ExecContext(ctx,
		seller.Nama,
		seller.Hp,
		seller.Kode,
		seller.Password,
		seller.Alamat,
		seller.Status,
		seller.RegionalId,
		dto.CurrUser,
		util.GetCurrTimeString(),
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

	currTime := util.GetCurrTimeString()
	db := database.GetConn()

	sqlStatement := `
		UPDATE sellers
		SET nama=?, last_update_by=?, last_update=?, hp=?, alamat=?, regional_id=?
		WHERE id=?;
	`

	res, err := db.Exec(
		sqlStatement,
		seller.Nama, dto.CurrUser, currTime, seller.Hp, seller.Alamat, seller.RegionalId, seller.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func UpdateStatusSeller(idSeller int64, statusSeller interface{}) (string, error) {

	currTime := util.GetCurrTimeString()
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

	currTime := util.GetCurrTimeString()
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

func generateQuerySeller(searchRequestDto dto.SearchRequestDto, page, limit int) (string, string) {

	sqlFind := `
		SELECT s.id, s.nama, s.hp, s.kode,  s.alamat, s.status, s.regional_id, s.last_update_by, s.last_update, r.nama 
		FROM sellers s	
		left join regional r on r.id = s.regional_id   
		WHERE s.nama like '%' and s.kode like '%' 
	`
	sqlCount := `
		SELECT count(*)
		FROM db_antar_barang.sellers	
		WHERE nama like '%' and kode like '%' 
	`
	return sqlFind, sqlCount

}

func GetSellerPage(searchRequestDto dto.SearchRequestDto, page, count int) ([]model.Seller, int, error) {
	db := database.GetConn()
	var sellers []model.Seller
	var total int

	sqlFind, sqlCount := generateQuerySeller(searchRequestDto, page, count)

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerySearchSeller(db, sqlFind, &sellers, errQuery)
	go AsyncQueryCount(db, sqlCount, &total, errCount)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return sellers, 0, resErrCount
	}

	if resErrQuery != nil {
		return sellers, 0, resErrQuery
	}

	return sellers, total, nil
}

func AsyncQuerySearchSeller(db *sql.DB, sqlFindSeller string, sellers *[]model.Seller, resChan chan error) {

	datas, err := db.QueryContext(
		context.Background(),
		sqlFindSeller)

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		resChan <- err
		return
	}

	for datas.Next() {
		var seller model.Seller
		err = datas.Scan(
			&seller.ID,
			&seller.Nama,
			&seller.Hp,
			&seller.Kode,
			&seller.Alamat,
			&seller.Status,
			&seller.RegionalId,
			&seller.LastUpdateBy,
			&seller.LastUpdate,
			&seller.RegionalName,
		)
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*sellers = append(*sellers, seller)
	}
	resChan <- nil

}
