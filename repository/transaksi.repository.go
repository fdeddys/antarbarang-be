package repository

import (
	"errors"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func NewTransaksi(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn

	var customer model.Customer

	errCustomer := db().
		QueryRow(`
			SELECT * 
			FROM public.cutomers
			WHERE id = $1
			`,
			transaksi.IdCustomer,
		).Scan(&customer)
	if errCustomer != nil {
		return transaksi, errors.New("Error Table Customer : " + errCustomer.Error())
	}

	lastInsertId := 0
	transaksi.CoordinateTujuan = customer.Coordinate
	transaksi.TransaksiDate = util.GetCurrTimeUnix()
	transaksi.Status = enumerate.StatusTransaksi(enumerate.NEW)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		INSERT INTO public.transaksi
			(transaksi_date, jam_request_antar, tanggal_request_antar, nama_product, status, coordinate_tujuan, keterangan, id_seller, id_customer, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

	err := db().
		QueryRow(
			sqlStatement,
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
		).
		Scan(&lastInsertId)

	if err != nil {
		return transaksi, err
	}
	transaksi.ID = int64(lastInsertId)
	return transaksi, nil
}

func OnProccess(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.ON_PROCCESS)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser

	sqlStatement := `
		UPDATE public.transaksi
		SET
			id_driver = $1,
			status = $2,
			last_update_by = $3,
			last_update = $4
		WHERE	
			id = $5
	`

	_, err := db.Exec(
		sqlStatement,
		transaksi.IdDriver,
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

func OnTheWay(transaksi model.Transaksi) (model.Transaksi, error) {
	db := database.GetConn()

	transaksi.Status = enumerate.StatusTransaksi(enumerate.ON_THE_WAY)
	transaksi.LastUpdate = util.GetCurrTimeUnix()
	transaksi.LastUpdateBy = dto.CurrUser
	transaksi.TanggalAmbil = util.GetCurrTimeUnix()

	sqlStatement := `
		UPDATE public.transaksi
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
