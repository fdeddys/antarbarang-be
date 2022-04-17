package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	// conn     *pgx.Conn
	db       *sql.DB
	host     string
	port     int32
	user     string
	password string
	dbname   string
)

func init() {

	host = os.Getenv("DATABASE_HOST")
	user = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASS")
	dbname = os.Getenv("DATABASE_DBNAME")
	var err error
	db, err = openDatabase()
	if err != nil {
		fmt.Println("Failed Open Database!")
	}
}

func openDatabase() (*sql.DB, error) {
	port, errPort := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 32)
	if errPort != nil {
		port = 5432
	}
	fmt.Println("User : ", user, " Password : ", password, " host :", host, "port : ", port, "dbname ", dbname)
	dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?tls=skip-verify", user, password, host, port, dbname)
	//"user:password@/dbname"
	fmt.Println("Query string : ", dbUrl)
	var err error
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping database !", err.Error())
		return nil, err
	}
	// defer db.Close()

	fmt.Println("Database Successfully connected !")
	return db, nil
}

func GetConn() *sql.DB {
	if err := db.Ping(); err != nil {
		fmt.Println("Error ping database !", err.Error())
	}
	return db
}
