package repository

import (
	"errors"
	"fmt"

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
		SELECT id, nama, hp, alamat, status, last_update_by, last_update
		FROM public.sellers
		WHERE id = $1;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, id).
		Scan(
			&seller.ID,
			&seller.Nama,
			&seller.Hp,
			&seller.Alamat,
			&seller.Status,
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
		SELECT id, nama, hp, alamat, status, last_update_by, last_update
		FROM public.sellers
		WHERE kode = $1;
	`
	var seller model.Seller
	err := db.
		QueryRow(sqlStatement, kode).
		Scan(
			&seller.ID,
			&seller.Nama,
			&seller.Hp,
			&seller.Alamat,
			&seller.Status,
			&seller.LastUpdateBy,
			&seller.LastUpdate,
		)
	seller.LastUpdateStr = util.DateUnixToString(seller.LastUpdate)
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func SaveSeller(seller model.Seller) (model.Seller, error) {

	lastInsertId := 0
	db := database.GetConn
	kode, errKode := generateKodeSeller(constanta.PREFIX_SELLER)
	seller.Kode = kode
	seller.Status = enumerate.ACTIVE
	if errKode != nil {
		return model.Seller{}, errors.New("Cannot generate prefix cause : " + errKode.Error())
	}

	sqlStatement := `
		INSERT INTO public.sellers
			(nama, hp, kode, password, alamat,  status, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db().
		QueryRow(
			sqlStatement,
			seller.Nama, seller.Hp, seller.Kode, seller.Password, seller.Alamat, seller.Status, dto.CurrUser, util.GetCurrTimeUnix(),
		).
		Scan(&lastInsertId)

	if err != nil {
		return seller, err
	}
	seller.ID = int64(lastInsertId)
	return seller, nil
}

func generateKodeSeller(prefix string) (string, error) {
	// var urut model.Urut
	db := database.GetConn()
	// defer db.Close()

	row, err := db.
		Query(
			`SELECT no_terakhir from uruts where prefix = $1 FOR UPDATE `,
			prefix)
	if err != nil {
		fmt.Println("Error => ", err.Error())
		return "0", err
	}
	row.Next()

	var lastnumb int64
	err = row.Scan(&lastnumb)

	if err != nil {
		fmt.Println("error get value ", err)
	}
	nextNumb := lastnumb + 1

	// example
	// 0000099
	// index [ (len(string)-5), len(string)-1]
	// index [  (7-5)  , 7-1 ]
	// index [ 2, 6 ]

	// 000009
	// [(6-5) , (6-1))]
	// [1 , 5 ]

	// 00009
	// [(5-5), 5-1]
	// [0:4]
	// var result string
	result := fmt.Sprintf("%v%v", "00000", nextNumb)

	newKode := prefix + result[len(result)-5:]
	// len(result)-1
	fmt.Println("kode baru ", newKode)
	_, errUpd := db.Exec("UPDATE uruts set no_terakhir = $1 where prefix = $2", nextNumb, prefix)
	if errUpd != nil {
		return "0", err
	}
	// fmt.Println("Update urut : ", res.RowsAffected())
	return newKode, nil
}
