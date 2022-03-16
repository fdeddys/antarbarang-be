package repository

import (
	"context"
	"errors"
	"fmt"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindSellerByCode(id int) (model.Seller, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, hp, alamat, status, last_update_by, last_update
		FROM public.sellers
		WHERE kode = $1;
	`
	var seller model.Seller
	err := db().
		QueryRow(context.Background(), sqlStatement, id).
		Scan(&seller)
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
			context.Background(),
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
	fmt.Println("Prefix : ", prefix)
	// var urut model.Urut
	db := database.GetConn
	fmt.Println("68")
	row, err := db().
		Query(
			context.Background(),
			`SELECT id, prefix, keterangan, no_terakhir from uruts where prefix = $1; `,
			prefix)
	fmt.Println("74")
	if err != nil {
		fmt.Println("Error => ", err.Error())
		return "0", err
	}
	fmt.Println("79")
	row.Next()
	fmt.Println("81")
	data, err := row.Values()
	fmt.Println("83")
	if err != nil {
		fmt.Println("error get value ", err)
	}
	fmt.Println("87")
	nextNumb := data[3].(int64) + 1

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
	db().Exec(context.Background(), "UPDATE uruts set no_terakhir = no_terakhir +1 where prefix = $1", prefix)

	return newKode, nil
}
