package repository

import (
	"context"
	"database/sql"
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
		customer.Alamat,
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
			&customer.Alamat,
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
		customer.Nama, dto.CurrUser, currTime, customer.Hp, customer.Alamat, customer.Coordinate, customer.ID)

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

func generateQueryCustomer(searchRequestDto dto.SearchRequestDto, page, limit int) (string, string) {

	var kriteriaKode = "%"
	if searchRequestDto.Kode != "" {
		kriteriaKode += searchRequestDto.Kode + "%"
	}

	searchSellerID := false
	sellerId := int64(0)
	if len(searchRequestDto.SellerId) > 0 {
		searchSellerID = true
		sellerId, _ = strconv.ParseInt(searchRequestDto.SellerId, 10, 64)
	}

	sqlFind := `
		SELECT c.id, c.seller_id , 
			case when c.seller_id IS NULL or c.seller_id = ''
				then ' '
				else s.nama 
			end as seller ,
			c.nama, c.hp, c.alamat, c.coordinate, c.status, c.last_update_by, c.last_update
		FROM customers c `
	where := `
		left join sellers s on c.seller_id = s.id		
		WHERE ( c.nama like '%v' ) and ( ( not %v ) or (seller_id = %v) ) 
		ORDER BY c.nama  
	`
	limitQuery := `
		LIMIT %v, %v
	`

	sqlFind = fmt.Sprintf(sqlFind+where+limitQuery, kriteriaKode, searchSellerID, sellerId, ((page - 1) * limit), limit)
	fmt.Println("Query Find = ", sqlFind)

	sqlCount := `
		SELECT count(*)
		FROM customers c `
	sqlCount = fmt.Sprintf(sqlCount+where, kriteriaKode, searchSellerID, sellerId)
	fmt.Println("Query Count = ", sqlCount)

	return sqlFind, sqlCount

}

func GetCustomerPage(searchRequestDto dto.SearchRequestDto, page, count int) ([]model.Customer, int, error) {
	db := database.GetConn()
	var customers []model.Customer
	var total int

	sqlFind, sqlCount := generateQueryCustomer(searchRequestDto, page, count)

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go AsyncQuerySearchCustomer(db, sqlFind, &customers, errQuery)
	go AsyncQueryCount(db, sqlCount, &total, errCount)

	resErrCount := <-errCount
	resErrQuery := <-errQuery

	wg.Done()

	if resErrCount != nil {
		log.Println("errr-->", resErrCount)
		return customers, 0, resErrCount
	}

	if resErrQuery != nil {
		return customers, 0, resErrQuery
	}

	return customers, total, nil
}

func AsyncQuerySearchCustomer(db *sql.DB, sqlFind string, customers *[]model.Customer, resChan chan error) {

	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	if err != nil {
		fmt.Println("Error query context ", err.Error())
		resChan <- err
		return
	}

	for datas.Next() {
		var customer model.Customer
		err = datas.Scan(
			&customer.ID,
			&customer.SellerId,
			&customer.SellerName,
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
			fmt.Println("Error fetch data next : ", err.Error())
		}
		*customers = append(*customers, customer)
	}
	resChan <- nil
	return

}
