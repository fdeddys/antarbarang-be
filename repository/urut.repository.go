package repository

import (
	"fmt"

	"com.ddabadi.antarbarang/database"
)

func generateKode(prefix string) (string, error) {
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
