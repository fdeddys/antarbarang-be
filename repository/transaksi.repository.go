package repository

import (
	"context"
	"errors"
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
			id_driver = $1,
			id_admin = $2,
			status = $3,
			last_update_by = $4,
			last_update = $5
		WHERE	
			id = $6
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
			tanggal_ambil = $1,
			photo_ambil = $2,
			status = $3,
			last_update_by = $4,
			last_update = $5
		WHERE	
			id = $6
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
			tanggal_sampai = $1,
			photo_sampai = $2,
			status = $3,
			last_update_by = $4,
			last_update = $5
		WHERE	
			id = $6
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
