package repository

import (
	"context"
	"fmt"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindCustomerById(id int) (model.Customer, error) {
	db := database.GetConn

	sqlStatement := `
		SELECT id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM public.customers	
		WHERE id = $1;
	`
	var customer model.Customer
	err := db().
		QueryRow(context.Background(), sqlStatement, id).
		Scan(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil

}

func SaveCustomer(customer model.Customer) (int64, error) {

	lastInsertId := 0

	db := database.GetConn

	sqlStatement := `
		INSERT INTO public.customers
			(nama, hp, alamat, coordinate, status, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := db().
		QueryRow(context.Background(),
			sqlStatement, customer.Nama, customer.Hp, customer.Address, customer.Coordinate, enumerate.ACTIVE, dto.CurrUser, util.GetCurrTimeUnix()).
		Scan(&lastInsertId)

	if err != nil {
		return int64(lastInsertId), err
	}
	return int64(lastInsertId), nil

}

func FindCustomerByNama(nama string) ([]model.Customer, error) {

	customers := []model.Customer{}
	db := database.GetConn

	rows, err := db().Query(context.Background(), `
		SELECT id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE nama like '%$1%' `, nama)

	if err != nil {
		return customers, err
	}

	for rows.Next() {
		data, err := rows.Values()
		if err != nil {
			fmt.Println("Gagal ambil data dari query")
		}
		var customer model.Customer
		// convert DB types to Go types
		customer.Nama = data[0].(string)
		customer.Hp = data[1].(string)
		customer.Address = data[2].(string)
		customer.Coordinate = data[3].(string)
		customer.Status = data[4].(enumerate.StatusRecord)
		customer.LastUpdateBy = data[5].(string)
		customer.LastUpdate = data[6].(int64)

		customers = append(customers, customer)
	}

	return customers, nil

}
