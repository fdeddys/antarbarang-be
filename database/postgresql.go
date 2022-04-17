package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "adminap"
// 	password = "123"
// 	dbname   = "antarBarang"
// )

var (
// conn     *pgx.Conn
// db *sql.DB
// host     string
// port     int32
// user     string
// password string
// dbname   string
)

func initX() {

	host = os.Getenv("DATABASE_HOST")
	user = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASS")
	dbname = os.Getenv("DATABASE_DBNAME")
	if openDatabaseX() != nil {
		fmt.Println("Failed Open Database!")
	}
}

func openDatabaseX() error {
	port, errPort := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 32)
	if errPort != nil {
		port = 5432
	}
	fmt.Println("User : ", user, " Password : ", password, " host :", host, "port : ", port, "dbname ", dbname)
	dbUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Unable to parse DATABASE_URL", "error", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping database !", err.Error())
		return nil
	}

	fmt.Println("Database Successfully connected !")
	return nil
}

func GetConnX() *sql.DB {
	if err := db.Ping(); err != nil {
		fmt.Println("Error ping database !", err.Error())
	}
	return db
}

// func GetConn() *pgx.Conn {
// 	if err := conn.Ping(context.Background()); err != nil {
// 		fmt.Println("Error ping database !")
// 	}
// 	return conn
// }

// func openDatabasex() error {
// 	port, errPort := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 32)
// 	if errPort != nil {
// 		port = 5432
// 	}

// 	dbUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, password, host, port, dbname)
// 	// var err error
// 	// conn, err = pgx.Connect(context.Background(), dbUrl)

// 	// if err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
// 	// 	return err
// 	// }
// 	// logger := log15adapter.NewLogger(log.New("module", "pgx"))

// 	poolConfig, err := pgxpool.ParseConfig(dbUrl)
// 	if err != nil {
// 		// log.Crit("Unable to parse DATABASE_URL", "error", err)
// 		fmt.Println("Unable to parse DATABASE_URL", "error", err)
// 		os.Exit(1)
// 	}
// 	// logrusLogger := &logrus.Logger{
// 	// 	Out:          os.Stderr,
// 	// 	Formatter:    new(logrus.JSONFormatter),
// 	// 	Hooks:        make(logrus.LevelHooks),
// 	// 	Level:        logrus.InfoLevel,
// 	// 	ExitFunc:     os.Exit,
// 	// 	ReportCaller: false,
// 	//    }
// 	// poolConfig.ConnConfig.Logger = logrusLogger.Log(logrusLogger)

// 	// poolConfig.ConnConfig.Logger = logger
// 	poolConfig.ConnConfig.PreferSimpleProtocol = true

// 	// db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
// 	// if err != nil {
// 	// 	// log.Crit("Unable to create connection pool", "error", err)
// 	// 	os.Exit(1)
// 	// }
// 	fmt.Println("Database Successfully connected !")
// 	return nil
// }
