package repository

import (
	"fmt"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindCustomerById(id int) (model.Customer, error) {
	db := database.GetConn()
	defer db.Close()

	sqlStatement := `
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM public.customers	
		WHERE id = $1;
	`
	var customer model.Customer
	err := db.
		QueryRow(sqlStatement, id).
		Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.Nama,
			&customer.Hp,
			&customer.Alamat,
			&customer.Coordinate,
			&customer.Status,
			&customer.LastUpdateBy,
			&customer.LastUpdate,
		)
	customer.LastUpdateStr = util.DateUnixToString(customer.LastUpdate)
	if err != nil {
		return customer, err
	}
	return customer, nil

}

func SaveCustomer(customer model.Customer) (int64, error) {

	lastInsertId := int64(0)

	db := database.GetConn()
	defer db.Close()

	sqlStatement := `
		INSERT INTO public.customers
			(seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db.
		QueryRow(
			sqlStatement,
			customer.SellerId,
			customer.Nama,
			customer.Hp,
			customer.Alamat,
			customer.Coordinate,
			enumerate.ACTIVE,
			dto.CurrUser,
			util.GetCurrTimeUnix(),
		).
		Scan(&lastInsertId)
	// customer.LastUpdateStr = util.DateUnixToString(customer.LastUpdate)
	if err != nil {
		return lastInsertId, err
	}
	return lastInsertId, nil

}

func FindCustomerBySellerId(sellerId int64) ([]model.Customer, error) {

	customers := []model.Customer{}
	db := database.GetConn

	datas, err := db().Query(`
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE = $1 `,
		sellerId)

	if err != nil {
		return customers, err
	}
	for datas.Next() {
		var customer model.Customer
		err = datas.Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.Nama,
			&customer.Hp,
			&customer.Alamat,
			&customer.Coordinate,
			&customer.Status,
			&customer.LastUpdateBy,
			&customer.LastUpdate,
		)
		customer.LastUpdateStr = util.DateUnixToString(customer.LastUpdate)
		if err != nil {
			fmt.Println("Error fetch data next")
		}
		customers = append(customers, customer)
	}

	return customers, nil

}

func FindCustomerByNama(nama string) ([]model.Customer, error) {

	customers := []model.Customer{}
	db := database.GetConn

	datas, err := db().Query(`
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE nama like '%$1%' `,
		nama)

	if err != nil {
		return customers, err
	}
	for datas.Next() {
		var customer model.Customer
		err = datas.Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.Nama,
			&customer.Hp,
			&customer.Alamat,
			&customer.Coordinate,
			&customer.Status,
			&customer.LastUpdateBy,
			&customer.LastUpdate,
		)
		customer.LastUpdateStr = util.DateUnixToString(customer.LastUpdate)
		if err != nil {
			fmt.Println("Error fetch data next")
		}
		customers = append(customers, customer)
	}

	return customers, nil

}
