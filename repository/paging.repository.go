package repository

import "database/sql"

func AsyncQueryCount(db *sql.DB, sqlCount string, total *int, errCount chan error) {

	totalRec := 0
	err := db.
		QueryRow(sqlCount).
		Scan(
			&totalRec,
		)
	if err != nil {
		errCount <- err
		return
	}
	*total = totalRec
	errCount <- nil

}
