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

	var seller model.Seller
	sqlSearchSeller := ` 
		SELECT s.id , s.regional_id, rg.nama 
		FROM sellers s
		left join regional r  on s.regional_id = r.id 
		left join regional_group rg  on rg.id  = r.regional_group_id 
		WHERE s.id = ?
	`
	errSeller := db.QueryRow(
		sqlSearchSeller,
		transaksi.IdSeller,
	).Scan(&seller.ID, &seller.RegionalId, &seller.RegionalGroupName)

	if errSeller != nil {
		return transaksi, errors.New("Error Table Seller : " + errSeller.Error())
	}

	var customer model.Customer
	sqlSearchCustomer := ` 
		SELECT c.id , c.coordinate, c.regional_id, rg.nama 
		FROM customers c  
		left join regional r  on c.regional_id = r.id 
		left join regional_group rg  on rg.id  = r.regional_group_id 
		WHERE c.id = ?
	`
	errCustomer := db.QueryRow(
		sqlSearchCustomer,
		transaksi.IdCustomer,
	).Scan(&customer.ID, &customer.Coordinate, &customer.RegionalId, &customer.RegionalGroupName)

	if errCustomer != nil {
		return transaksi, errors.New("Error Table Customer : " + errCustomer.Error())
	}

	transaksi.RegionalSeller = seller.RegionalId
	transaksi.RegionalGroupSeller = seller.RegionalGroupName
	transaksi.CoordinateTujuan = customer.Coordinate
	transaksi.RegionalCustomer = customer.RegionalId
	transaksi.RegionalGroupCustomer = customer.RegionalGroupName
	transaksi.TransaksiDate = util.GetCurrDate().Format("2006-01-02")
	transaksi.Status = enumerate.StatusTransaksi(enumerate.NEW)
	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
	transaksi.LastUpdateBy = dto.CurrUser

	fmt.Println("Transaksi = ", transaksi)
	sqlStatement := `
		INSERT INTO transaksi
			(transaksi_date, jam_request_antar, tanggal_request_antar, nama_product, status, coordinate_tujuan, keterangan, id_seller, id_customer, last_update_by, last_update, regional_seller, regional_customer, regional_group_seller, regional_group_customer )
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
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
		transaksi.RegionalSeller,
		transaksi.RegionalCustomer,
		transaksi.RegionalGroupSeller,
		transaksi.RegionalGroupCustomer,
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
	transaksi.RegionalCustomer = customer.RegionalId
	transaksi.TransaksiDate = util.GetCurrDate().Format("2006-01-02")
	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
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

func OnProccessRepo(transaksi model.Transaksi) (int, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.ON_PROCCESS)
	transaksi.TanggalOnProccess = util.GetCurrDate().Format("2006-01-02 15:04:05")
	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		UPDATE transaksi
		SET
			id_driver = ?,
			status = ?,
			tanggal_on_proccess = ?,
			last_update_by = ?,
			last_update = ?
		WHERE	
			id = ?
	`

	res, err := db.Exec(
		sqlStatement,
		transaksi.IdDriver,
		transaksi.Status,
		transaksi.TanggalOnProccess,
		transaksi.LastUpdateBy,
		transaksi.LastUpdate,
		transaksi.ID,
	)

	totalRec, _ := res.RowsAffected()
	fmt.Println("res update =>", totalRec)
	if err != nil {
		return 0, errors.New("Error transaksi : " + err.Error())
	}

	return int(totalRec), nil
}

func UpdateOnProccessRepo(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
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
	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
	transaksi.LastUpdateBy = dto.CurrUser
	transaksi.TanggalAmbil = util.GetCurrDate().Format("2006-01-02 15:04:05")

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
	transaksi.LastUpdate = util.GetCurrDate().Format("2006-01-02 15:04:05")
	transaksi.LastUpdateBy = dto.CurrUser
	transaksi.TanggalSampai = util.GetCurrDate().Format("2006-01-02 15:04:05")

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
	// var kriteriaDriver = "%"
	// if searchTransaksiRequestDto.DriverName != "" {
	// 	kriteriaDriver += searchTransaksiRequestDto.DriverName + "%"
	// }
	var kriteriaCustomer = "%"
	if searchTransaksiRequestDto.CustomerName != "" {
		kriteriaCustomer += searchTransaksiRequestDto.CustomerName + "%"
	}
	// searchStatus := false
	// status := int64(0)
	// if len(searchTransaksiRequestDto.Status) > 0 {
	// 	searchStatus = true
	// 	status, _ = strconv.ParseInt(searchTransaksiRequestDto.Status, 10, 64)
	// }

	searchStatus := false
	status := "0"
	if len(searchTransaksiRequestDto.Status) > 0 {
		searchStatus = true
		status = searchTransaksiRequestDto.Status
	}

	searchDriverID := false
	driverID := int64(0)
	if len(searchTransaksiRequestDto.DriverID) > 0 {
		searchDriverID = true
		driverID, _ = strconv.ParseInt(searchTransaksiRequestDto.DriverID, 10, 64)
	}

	tgl1 := searchTransaksiRequestDto.Tgl1
	tgl2 := searchTransaksiRequestDto.Tgl2

	sqlFind := `
		SELECT  t.id, 
			transaksi_date, 
			tanggal_request_antar, 
			jam_request_antar, 
			nama_product, 
			t.status, 
			CASE
				WHEN t.status = 0 THEN "NEW"
				WHEN t.status = 1 THEN "ON_PROCCESS"
				WHEN t.status = 2 THEN "ON_THE_WAY"
				WHEN t.status = 3 THEN "DONE"
				WHEN t.status = 4 THEN "CANCEL"
				ELSE "UNKNOWN"
			END,
			coordinate_tujuan, 
			keterangan, 
			IFNULL(photo_ambil,""), 
			IFNULL(tanggal_ambil,0), 
			IFNULL(photo_sampai,""), 
			IFNULL(tanggal_sampai,0), 
			id_seller, 
			IFNULL(s.nama,"") , 
			IFNULL(s.alamat,"") , 
			IFNULL(s.hp,"") , 
			IFNULL(id_driver,0), 
			IFNULL(d.nama,"") , 
			id_customer, 
			IFNULL(c.nama,"") , 
			IFNULL(c.alamat,"") , 
			IFNULL(c.hp,"") , 
			IFNULL(id_admin,0), 
			t.last_update_by, 
			t.last_update,
			t.regional_seller,
			t.regional_group_seller,
			t.regional_customer ,
			t.regional_group_customer,
			c.lng,
			c.lat 
		FROM transaksi t
		left join sellers s on t.id_seller  = s.id
		left join drivers d on t.id_driver = d.id 
		left JOIN customers c on t.id_customer = c.id  `
	where := `
		WHERE ( c.nama like '%v' ) 
		AND   ( s.nama  like '%v' ) 
		AND	  ( ( not %v  ) or (t.status  in (%v) ) )
		AND	  ( ( not %v  ) or (t.id_driver  = %v ) )
		AND   (transaksi_date) BETWEEN  '%v 00:00:00' and  '%v 23:59:59'
		ORDER BY t.transaksi_date DESC   
	`
	limitQuery := `
		LIMIT %v, %v
	`
	// kriteriaDriver,
	sqlFind = fmt.Sprintf(
		sqlFind+where+limitQuery,
		kriteriaCustomer,
		kriteriaSeller,
		searchStatus,
		status,
		searchDriverID,
		driverID,
		tgl1,
		tgl2,
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
		searchStatus,
		status,
		searchDriverID,
		driverID,
		tgl1,
		tgl2,
	)
	// kriteriaDriver,
	fmt.Println("Query Count = ", sqlCount)

	return sqlFind, sqlCount

}

func GetTransaksiPage(searchRequestDto dto.SearchTransaksiRequestDto, page, count int) ([]model.Transaksi, int, error) {
	db := database.GetConn()
	// var transaksis []model.Transaksi = []model.Transaksi
	transaksis := make([]model.Transaksi, 0)
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

	fmt.Println("Query search : ", sqlFind)

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
			&transaksi.StatusName,
			&transaksi.CoordinateTujuan,
			&transaksi.Keterangan,
			&transaksi.PhotoAmbil,
			&transaksi.TanggalAmbil,
			&transaksi.PhotoSampai,
			&transaksi.TanggalSampai,
			&transaksi.IdSeller,
			&transaksi.SellerName,
			&transaksi.SellerAddress,
			&transaksi.SellerHp,
			&transaksi.IdDriver,
			&transaksi.DriverName,
			&transaksi.IdCustomer,
			&transaksi.CustomerName,
			&transaksi.CustomerAddress,
			&transaksi.CustomerHp,
			&transaksi.IdAdmin,
			&transaksi.LastUpdateBy,
			&transaksi.LastUpdate,
			&transaksi.RegionalSeller,
			&transaksi.RegionalGroupSeller,
			&transaksi.RegionalCustomer,
			&transaksi.RegionalGroupCustomer,
			&transaksi.CustLng,
			&transaksi.CustLat,
		)
		transaksi.TransaksiDateStr = transaksi.TransaksiDate
		transaksi.LastUpdateStr = (transaksi.LastUpdate)
		transaksi.TanggalAmbilStr = transaksi.TanggalAmbil
		transaksi.TanggalSampaiStr = transaksi.TanggalSampai
		transaksi.TanggalRequestAntarStr = transaksi.TanggalRequestAntar
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*transaksis = append(*transaksis, transaksi)
	}
	resChan <- nil
	return

}

func GetTransaksiByDriverByTanggalAntar(searchRequestDto dto.SearchTransaksiRequestDto) ([]model.Transaksi, error) {
	db := database.GetConn()
	transaksis := make([]model.Transaksi, 0)

	sqlFind := generateQueryTransaksiByDriverByTanggalAntar(searchRequestDto)

	fmt.Println("Query search : ", sqlFind)

	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	for datas.Next() {
		var transaksi model.Transaksi
		err = datas.Scan(
			&transaksi.ID,
			&transaksi.TransaksiDate,
			&transaksi.TanggalRequestAntar,
			&transaksi.JamRequestAntar,
			&transaksi.NamaProduct,
			&transaksi.Status,
			&transaksi.StatusName,
			&transaksi.CoordinateTujuan,
			&transaksi.Keterangan,
			&transaksi.PhotoAmbil,
			&transaksi.TanggalAmbil,
			&transaksi.PhotoSampai,
			&transaksi.TanggalSampai,
			&transaksi.IdSeller,
			&transaksi.SellerName,
			&transaksi.SellerAddress,
			&transaksi.SellerHp,
			&transaksi.IdDriver,
			&transaksi.DriverName,
			&transaksi.IdCustomer,
			&transaksi.CustomerName,
			&transaksi.CustomerAddress,
			&transaksi.CustomerHp,
			&transaksi.IdAdmin,
			&transaksi.LastUpdateBy,
			&transaksi.LastUpdate,
			&transaksi.CustLng,
			&transaksi.CustLat,
		)
		transaksi.TransaksiDateStr = transaksi.TransaksiDate
		transaksi.LastUpdateStr = (transaksi.LastUpdate)
		transaksi.TanggalAmbilStr = transaksi.TanggalAmbil
		transaksi.TanggalSampaiStr = transaksi.TanggalSampai
		transaksi.TanggalRequestAntarStr = transaksi.TanggalRequestAntar
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		transaksis = append(transaksis, transaksi)
	}

	return transaksis, nil
}

func generateQueryTransaksiByDriverByTanggalAntar(searchTransaksiRequestDto dto.SearchTransaksiRequestDto) string {

	var kriteriaSeller = "%"
	if searchTransaksiRequestDto.SellerName != "" {
		kriteriaSeller += searchTransaksiRequestDto.SellerName + "%"
	}

	searchStatus := false
	status := "0"
	if len(searchTransaksiRequestDto.Status) > 0 {
		searchStatus = true
		status = searchTransaksiRequestDto.Status
	}

	searchDriverID := false
	driverID := int64(0)
	if len(searchTransaksiRequestDto.DriverID) > 0 {
		searchDriverID = true
		driverID, _ = strconv.ParseInt(searchTransaksiRequestDto.DriverID, 10, 64)
	}

	tgl1 := searchTransaksiRequestDto.Tgl1
	tgl2 := searchTransaksiRequestDto.Tgl2

	sqlFind := `
		SELECT  t.id, 
			transaksi_date, 
			tanggal_request_antar, 
			jam_request_antar, 
			nama_product, 
			t.status, 
			CASE
				WHEN t.status = 0 THEN "NEW"
				WHEN t.status = 1 THEN "ON_PROCCESS"
				WHEN t.status = 2 THEN "ON_THE_WAY"
				WHEN t.status = 3 THEN "DONE"
				WHEN t.status = 4 THEN "CANCEL"
				ELSE "UNKNOWN"
			END,
			coordinate_tujuan, 
			keterangan, 
			IFNULL(photo_ambil,""), 
			IFNULL(tanggal_ambil,0), 
			IFNULL(photo_sampai,""), 
			IFNULL(tanggal_sampai,0), 
			id_seller, 
			IFNULL(s.nama,"") , 
			IFNULL(s.alamat,"") , 
			IFNULL(s.hp,"") , 
			IFNULL(id_driver,0), 
			IFNULL(d.nama,"") , 
			id_customer, 
			IFNULL(c.nama,"") , 
			IFNULL(c.alamat,"") , 
			IFNULL(c.hp,"") , 
			IFNULL(id_admin,0), 
			t.last_update_by, 
			t.last_update,
			c.lng,
			c.lat
		FROM transaksi t
		left join sellers s on t.id_seller  = s.id
		left join drivers d on t.id_driver = d.id 
		left JOIN customers c on t.id_customer = c.id  `
	where := `
		WHERE  ( s.nama  like '%v' ) 
		AND ( ( not %v  ) or (t.status  in (%v) ) )
		AND	  ( ( not %v  ) or (t.id_driver  = %v ) )
		AND   (tanggal_request_antar) BETWEEN  '%v 00:00:00' and  '%v 23:59:59'
		ORDER BY t.jam_request_antar  ASC   
	`

	// kriteriaDriver,
	sqlFind = fmt.Sprintf(
		sqlFind+where,
		kriteriaSeller,
		searchStatus,
		status,
		searchDriverID,
		driverID,
		tgl1,
		tgl2,
	)
	fmt.Println("Query Find = ", sqlFind)

	return sqlFind

}

func GetTransaksiByTglAntarPage(searchRequestDto dto.SearchTransaksiRequestDto, page, count int) ([]model.Transaksi, int, error) {
	db := database.GetConn()
	transaksis := make([]model.Transaksi, 0)
	var total int

	sqlFind, sqlCount := generateQueryTransaksiPickupByTglAntar(searchRequestDto, page, count)

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

func generateQueryTransaksiPickupByTglAntar(searchTransaksiRequestDto dto.SearchTransaksiRequestDto, page, limit int) (string, string) {

	var kriteriaSeller = "%"
	if searchTransaksiRequestDto.SellerName != "" {
		kriteriaSeller += searchTransaksiRequestDto.SellerName + "%"
	}

	var kriteriaCustomer = "%"
	if searchTransaksiRequestDto.CustomerName != "" {
		kriteriaCustomer += searchTransaksiRequestDto.CustomerName + "%"
	}

	searchStatus := false
	status := "0"
	if len(searchTransaksiRequestDto.Status) > 0 {
		searchStatus = true
		status = searchTransaksiRequestDto.Status
	}

	searchDriverID := false
	driverID := int64(0)
	if len(searchTransaksiRequestDto.DriverID) > 0 {
		searchDriverID = true
		driverID, _ = strconv.ParseInt(searchTransaksiRequestDto.DriverID, 10, 64)
	}

	tgl1 := searchTransaksiRequestDto.Tgl1
	tgl2 := searchTransaksiRequestDto.Tgl2

	sqlFind := `
		SELECT  t.id, 
			transaksi_date, 
			tanggal_request_antar, 
			jam_request_antar, 
			nama_product, 
			t.status, 
			CASE
				WHEN t.status = 0 THEN "NEW"
				WHEN t.status = 1 THEN "ON_PROCCESS"
				WHEN t.status = 2 THEN "ON_THE_WAY"
				WHEN t.status = 3 THEN "DONE"
				WHEN t.status = 4 THEN "CANCEL"
				ELSE "UNKNOWN"
			END,
			coordinate_tujuan, 
			keterangan, 
			IFNULL(photo_ambil,""), 
			IFNULL(tanggal_ambil,0), 
			IFNULL(photo_sampai,""), 
			IFNULL(tanggal_sampai,0), 
			id_seller, 
			IFNULL(s.nama,"") , 
			IFNULL(s.alamat,"") , 
			IFNULL(s.hp,"") , 
			IFNULL(id_driver,0), 
			IFNULL(d.nama,"") , 
			id_customer, 
			IFNULL(c.nama,"") , 
			IFNULL(c.alamat,"") , 
			IFNULL(c.hp,"") , 
			IFNULL(id_admin,0), 
			t.last_update_by, 
			t.last_update,
			IFNULL(t.regional_seller,0),
			IFNULL(t.regional_group_seller,""),
			IFNULL(t.regional_customer,0) ,
			IFNULL(t.regional_group_customer,""),
			c.lng,
			c.lat
		FROM transaksi t
		left join sellers s on t.id_seller  = s.id
		left join drivers d on t.id_driver = d.id 
		left JOIN customers c on t.id_customer = c.id  `
	where := `
		WHERE ( c.nama like '%v' ) 
		AND   ( s.nama  like '%v' ) 
		AND	  ( ( not %v  ) or (t.status  in (%v) ) )
		AND	  ( ( not %v  ) or (t.id_driver  = %v ) )
		AND   (tanggal_request_antar) BETWEEN  '%v 00:00:00' and  '%v 23:59:59'
		ORDER BY t.tanggal_request_antar DESC   
	`
	limitQuery := `
		LIMIT %v, %v
	`
	// kriteriaDriver,
	sqlFind = fmt.Sprintf(
		sqlFind+where+limitQuery,
		kriteriaCustomer,
		kriteriaSeller,
		searchStatus,
		status,
		searchDriverID,
		driverID,
		tgl1,
		tgl2,
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
		searchStatus,
		status,
		searchDriverID,
		driverID,
		tgl1,
		tgl2,
	)
	// kriteriaDriver,
	fmt.Println("Query Count = ", sqlCount)

	return sqlFind, sqlCount

}
