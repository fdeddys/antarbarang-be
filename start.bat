set PORT=8888
set DATABASE_HOST=156.67.214.228
set DATABASE_USER=admin
set DATABASE_PASS=Pr@MoT10n
set DATABASE_DBNAME=db_antar_barang
set DATABASE_PORT=3306
set REDIS_HOST=localhost
set REDIS_PORT=6379

go run main.go

@REM env GOOS=linux GOARCH=amd64 go build -o antar_be
