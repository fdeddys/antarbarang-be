package repository

import (
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
		SELECT id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM public.customers	
		WHERE id = $1;
	`
	var customer model.Customer
	err := db.
		QueryRow(sqlStatement, id).
		Scan(&customer)
	if err != nil {
		return customer, err
	}
	return customer, nil

}

func SaveCustomer(customer model.Customer) (int64, error) {

	lastInsertId := 0

	db := database.GetConn()
	defer db.Close()

	sqlStatement := `
		INSERT INTO public.customers
			(nama, hp, alamat, coordinate, status, last_update_by, last_update)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := db.
		QueryRow(
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

	rows, err := db().Query(`
		SELECT id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE nama like '%$1%' `, nama)

	if err != nil {
		return customers, err
	}

	for rows.Next() {

		var customer model.Customer
		err = rows.Scan(
			customer.Nama,
			customer.Hp,
			customer.Address,
			customer.Coordinate,
			customer.Status,
			customer.LastUpdateBy,
			customer.LastUpdate)
		if err != nil {
			return []model.Customer{}, err
		}

		customers = append(customers, customer)
	}

	return customers, nil

}
