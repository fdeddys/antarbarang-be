package repository

import (
	"context"
	"fmt"
	"time"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/util"
)

func FindCustomerById(id int) (model.Customer, error) {
	db := database.GetConn()
	// defer db.Close()

	sqlStatement := `
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE id = ?;
	`
	var customer model.Customer
	err := db.
		QueryRow(sqlStatement, id).
		Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.Nama,
			&customer.Hp,
			&customer.Address,
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
	// defer db.Close()

	sqlStatement := `
		INSERT INTO customers
			(seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(
		ctx,
		customer.SellerId,
		customer.Nama,
		customer.Hp,
		customer.Address,
		customer.Coordinate,
		enumerate.ACTIVE,
		dto.CurrUser,
		util.GetCurrTimeUnix(),
	)
	if err != nil {
		return 0, err
	}
	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func FindCustomerBySellerId(sellerId int64) ([]model.Customer, error) {

	customers := []model.Customer{}
	db := database.GetConn

	datas, err := db().Query(`
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE seller_id= ? `,
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
			&customer.Address,
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
	db := database.GetConn()

	sqlFindCustomerByNama := `
		SELECT id, seller_id, nama, hp, alamat, coordinate, status, last_update_by, last_update
		FROM customers	
		WHERE nama like ? 
	`
	//

	// var customer model.Customer
	datas, err := db.QueryContext(
		context.Background(),
		sqlFindCustomerByNama,
		"%"+nama+"%")
	// 	.
	// Scan(

	// )

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		return customers, err
	}

	for datas.Next() {
		var customer model.Customer
		err = datas.Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.Nama,
			&customer.Hp,
			&customer.Address,
			&customer.Coordinate,
			&customer.Status,
			&customer.LastUpdateBy,
			&customer.LastUpdate,
		)
		customer.LastUpdateStr = util.DateUnixToString(customer.LastUpdate)
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		customers = append(customers, customer)
	}

	return customers, nil

}

func UpdateCustomer(customer model.Customer) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE customers
		SET nama=?,  last_update_by=?, last_update=?, hp=?, alamat=?, coordinate = ?
		WHERE id=?;
	`

	res, err := db.Exec(
		sqlStatement,
		customer.Nama, dto.CurrUser, currTime, customer.Hp, customer.Address, customer.Coordinate, customer.ID)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}

func UpdateStatusCustomer(idCustomer int64, statusCustomer interface{}) (string, error) {

	currTime := util.GetCurrTimeUnix()
	db := database.GetConn()

	sqlStatement := `
		UPDATE customers
		SET status= ?,  last_update_by= ?, last_update= ?
		WHERE id  = ?;
	`

	res, err := db.Exec(
		sqlStatement,
		statusCustomer, dto.CurrUser, currTime, idCustomer)

	if err != nil {
		return "", err
	}
	totalData, _ := res.RowsAffected()
	return fmt.Sprintf("update data success : %v record's!", totalData), err
}
