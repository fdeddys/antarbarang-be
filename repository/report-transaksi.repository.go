package repository

import (
	"context"
	"fmt"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/dto"
)

func ReportTransaksiByDate(dateStart, dateEnd string) []dto.ReportTransaksi {
	db := database.GetConn()

	var dataTrx []dto.ReportTransaksi

	sqlFind := `
		SELECT t.tanggal_request_antar, t.keterangan , t.biaya, 
			IFNULL(t.regional_group_seller,""), t.id_seller , s.nama as 'seller',
			IFNULL(t.regional_group_customer,""), t.id_customer  , c.nama as 'cust',
			IFNULL(t.id_driver,0), 
			IFNULL(d.nama,"") as 'driver'
		FROM transaksi t
		LEFT JOIN customers c on c.id  = t.id_customer
		LEFT JOIN sellers s on s.id = t.id_seller 
		LEFT JOIN drivers d on d.id = t.id_driver 
		WHERE  (t.tanggal_request_antar) BETWEEN  '%v 00:00:00' and  '%v 23:59:59'
		ORDER BY t.tanggal_request_antar desc, s.nama asc
	`

	sqlFind = fmt.Sprintf(
		sqlFind,
		dateStart,
		dateEnd)

	fmt.Println("sql : ", sqlFind)
	datas, err := db.QueryContext(
		context.Background(),
		sqlFind)

	for datas.Next() {
		var transaksi dto.ReportTransaksi
		err = datas.Scan(
			&transaksi.TanggalAntar,
			&transaksi.Keterangan,
			&transaksi.Biaya,
			&transaksi.SellerRegion,
			&transaksi.SellerId,
			&transaksi.SellerName,
			&transaksi.CustomerRegion,
			&transaksi.CustomerId,
			&transaksi.CustomerName,
			&transaksi.DriverId,
			&transaksi.DriverName,
		)
		if err != nil {
			fmt.Println("Error fetch data next : ", err.Error())
		}
		dataTrx = append(dataTrx, transaksi)
	}
	fmt.Println("Total Record=", len(dataTrx))
	return dataTrx

}
