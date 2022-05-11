package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func NewTransaksiRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	var customer model.Customer
	sqlSearchCustomer := ` 
		SELECT id , coordinate
		FROM customers
		WHERE id = ?
	`
	errCustomer := db.QueryRow(
		sqlSearchCustomer,
		transaksi.IdCustomer,
	).Scan(&customer.ID, &customer.Coordinate)

	if errCustomer != nil {
		return transaksi, errors.New("Error Table Customer : " + errCustomer.Error())
	}

	transaksi.CoordinateTujuan = customer.Coordinate
	transaksi.TransaksiDate = util.GetCurrTimeUnix()
	transaksi.Status = enumerate.StatusTransaksi(enumerate.NEW)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		INSERT INTO transaksi
			(transaksi_date, jam_request_antar, tanggal_request_antar, nama_product, status, coordinate_tujuan, keterangan, id_seller, id_customer, last_update_by, last_update)
		VALUES (?,?,?,?,?,?,?,?,?,?,?)
	`
	ctx, errFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer errFunc()

	stmt, err := db.PrepareContext(
		ctx,
		sqlStatement,
	)
	if err != nil {
		return model.Transaksi{}, err
	}

	resp, err := stmt.ExecContext(
		ctx,
		transaksi.TransaksiDate,
		transaksi.JamRequestAntar,
		transaksi.TanggalRequestAntar,
		transaksi.NamaProduct,
		transaksi.Status,
		transaksi.CoordinateTujuan,
		transaksi.Keterangan,
		transaksi.IdSeller,
		transaksi.IdCustomer,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
	)

	if err != nil {
		return model.Transaksi{}, err
	}
	idTrx, err := resp.LastInsertId()
	if err != nil {
		return model.Transaksi{}, err
	}

	transaksi.ID = idTrx
	// err := db().
	// 	QueryRow(
	// 		sqlStatement,

	// 	).
	// 	Scan(&lastInsertId)

	// if err != nil {
	// 	return transaksi, err
	// }
	// transaksi.ID = int64(lastInsertId)
	return transaksi, nil
}

func UpdateNewTransaksiRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn

	var customer model.Customer

	errCustomer := db().
		QueryRow(`
				SELECT * 
				FROM cutomers
				WHERE id = $1
			`,
			transaksi.IdCustomer,
		).Scan(&customer)
	if errCustomer != nil {
		return transaksi, errors.New("Error Table Customer : " + errCustomer.Error())
	}

	transaksi.CoordinateTujuan = customer.Coordinate
	transaksi.TransaksiDate = util.GetCurrTimeUnix()
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		UPDATE transaksi
		SET
			jam_request_antar = $1, 
			tanggal_request_antar = $2, 
			nama_product = $3, 
			coordinate_tujuan = $4, 
			keterangan = $5, 
			id_customer = $6, 
			last_update_by = $7, 
			last_update = $8
		WHERE id = $9
		`

	_, err := db().
		Exec(
			sqlStatement,
			transaksi.JamRequestAntar,
			transaksi.TanggalRequestAntar,
			transaksi.NamaProduct,
			transaksi.CoordinateTujuan,
			transaksi.Keterangan,
			transaksi.IdCustomer,
			transaksi.LastUpdateBy,
			transaksi.LastUpdate,
		)

	if err != nil {
		return transaksi, err
	}
	return transaksi, nil

}

func OnProccessRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.ON_PROCCESS)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		UPDATE transaksi
		SET
			id_driver = ?,
			id_admin = ?,
			status = ?,
			last_update_by = ?,
			last_update = ?
		WHERE	
			id = ?
	`

	_, err := db.Exec(
		sqlStatement,
		transaksi.IdDriver,
		transaksi.IdAdmin,
		transaksi.Status,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
		transaksi.ID,
	)

	if err != nil {
		return transaksi, errors.New("Error transaksi : " + err.Error())
	}

	return transaksi, nil
}

func UpdateOnProccessRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		UPDATE transaksi
		SET
			id_driver = $1,
			id_admin = $2,
			last_update_by = $3,
			last_update = $4
		WHERE	
			id = $5
	`

	_, err := db.Exec(
		sqlStatement,
		transaksi.IdDriver,
		transaksi.IdAdmin,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
		transaksi.ID,
	)

	if err != nil {
		return transaksi, errors.New("Error transaksi : " + err.Error())
	}

	return transaksi, nil
}

func OnTheWayRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.ON_THE_WAY)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser
	transaksi.TanggalAmbil = util.GetCurrTimeUnix()

	sqlStatement := `
		UPDATE transaksi
		SET
			tanggal_ambil = ?,
			photo_ambil = ?,
			status = ?,
			last_update_by = ?,
			last_update = ?
		WHERE	
			id = ?
	`

	_, err := db.Exec(
		sqlStatement,
		transaksi.TanggalAmbil,
		transaksi.PhotoAmbil,
		transaksi.Status,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
		transaksi.ID,
	)

	if err != nil {
		return transaksi, errors.New("Error transaksi : " + err.Error())
	}

	return transaksi, nil
}

func DoneRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.DONE)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser
	transaksi.TanggalSampai = util.GetCurrTimeUnix()

	sqlStatement := `
		UPDATE transaksi
		SET
			tanggal_sampai = ?,
			photo_sampai = ?,
			status = ?,
			last_update_by = ?,
			last_update = ?
		WHERE	
			id = ?
	`

	_, err := db.Exec(
		sqlStatement,
		transaksi.TanggalSampai,
		transaksi.PhotoSampai,
		transaksi.Status,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
		transaksi.ID,
	)

	if err != nil {
		return transaksi, errors.New("Error transaksi : " + err.Error())
	}

	return transaksi, nil
}

func generateQueryTransaksi(searchTransaksiRequestDto dto.SearchTransaksiRequestDto, page, limit int) (string, string) {

	var kriteriaSeller = "%"
	if searchTransaksiRequestDto.SellerName != "" {
		kriteriaSeller += searchTransaksiRequestDto.SellerName + "%"
	}
	var kriteriaDriver = "%"
	if searchTransaksiRequestDto.DriverName != "" {
		kriteriaDriver += searchTransaksiRequestDto.DriverName + "%"
	}
	var kriteriaCustomer = "%"
	if searchTransaksiRequestDto.CustomerName != "" {
		kriteriaCustomer += searchTransaksiRequestDto.CustomerName + "%"
	}
	searchStatus := false
	status := int64(0)
	if len(searchTransaksiRequestDto.Status) > 0 {
		searchStatus = true
		status, _ = strconv.ParseInt(searchTransaksiRequestDto.Status, 10, 64)
	}

	sqlFind := `
		SELECT  t.id, 
			transaksi_date, 
			tanggal_request_antar, 
			jam_request_antar, 
			nama_product, 
			t.status, 
			coordinate_tujuan, 
			keterangan, 
			photo_ambil, 
			tanggal_ambil, 
			photo_sampai, 
			tanggal_sampai, 
			id_seller, 
			s.nama , 
			id_driver, 
			d.nama , 
			id_customer, 
			c.nama , 
			id_admin, 
			t.last_update_by, 
			t.last_update
		FROM transaksi t
		left join sellers s on t.id_seller  = s.id
		left join drivers d on t.id_driver = d.id 
		left JOIN customers c on t.id_customer = c.id  `
	where := `
		WHERE ( ( c.nama like '%v' ) OR ( s.nama  like '%v' ) OR ( d.nama  like '%v' ) )
		AND	  ( ( not %v  ) or (t.status  = %v ) )
		ORDER BY t.transaksi_date DESC   
	`
	limitQuery := `
		LIMIT %v, %v
	`

	sqlFind = fmt.Sprintf(
		sqlFind+where+limitQuery,
		kriteriaCustomer,
		kriteriaSeller,
		kriteriaDriver,
		searchStatus,
		status,
		((page - 1) * limit), limit)
	fmt.Println("Query Find = ", sqlFind)

	sqlCount := `
		SELECT count(*)
		FROM transaksi t
		left join sellers s on t.id_seller  = s.id
		left join drivers d on t.id_driver = d.id 
		left JOIN customers c on t.id_customer = c.id   `
	sqlCount = fmt.Sprintf(
		sqlCount+where,
		kriteriaCustomer,
		kriteriaSeller,
		kriteriaDriver,
		searchStatus,
		status,
	)
	fmt.Println("Query Count = ", sqlCount)

	return sqlFind, sqlCount

}

func GetTransaksiPage(searchRequestDto dto.SearchTransaksiRequestDto, page, count int) ([]model.Transaksi, int, error) {
	db := database.GetConn()
	var transaksis []model.Transaksi
	var total int

	sqlFind, sqlCount := generateQueryTransaksi(searchRequestDto, page, count)

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerySearchTransaksi(db, sqlFind, &transaksis, errQuery)
	go AsyncQueryCount(db, sqlCount, &total, errCount)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return transaksis, 0, resErrCount
	}

	if resErrQuery != nil {
		return transaksis, 0, resErrQuery
	}

	return transaksis, total, nil
}

func AsyncQuerySearchTransaksi(db *sql.DB, sqlFind string, transaksis *[]model.Transaksi, resChan chan error) {

	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		resChan <- err
		return
	}

	for datas.Next() {
		var transaksi model.Transaksi
		err = datas.Scan(
			&transaksi.ID,
			&transaksi.TransaksiDate,
			&transaksi.TanggalRequestAntar,
			&transaksi.JamRequestAntar,
			&transaksi.NamaProduct,
			&transaksi.Status,
			&transaksi.CoordinateTujuan,
			&transaksi.Keterangan,
			&transaksi.PhotoAmbil,
			&transaksi.TanggalAmbil,
			&transaksi.PhotoSampai,
			&transaksi.TanggalSampai,
			&transaksi.IdSeller,
			&transaksi.SellerName,
			&transaksi.IdDriver,
			&transaksi.DriverName,
			&transaksi.IdCustomer,
			&transaksi.CustomerName,
			&transaksi.IdAdmin,
			&transaksi.LastUpdateBy,
			&transaksi.LastUpdate,
		)
		transaksi.TransaksiDateStr = util.DateUnixToString(transaksi.TransaksiDate)
		transaksi.LastUpdateStr = util.DateUnixToString(transaksi.LastUpdate)
		transaksi.TanggalAmbilStr = util.DateUnixToString(transaksi.TanggalAmbil)
		transaksi.TanggalSampaiStr = util.DateUnixToString(transaksi.TanggalSampai)
		transaksi.TanggalRequestAntarStr = util.DateUnixToString(transaksi.TanggalRequestAntar)
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*transaksis = append(*transaksis, transaksi)
	}
	resChan <- nil
	return

}
