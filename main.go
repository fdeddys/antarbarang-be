package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"

	"com.ddabadi.antarbarang/database"
	_ "com.ddabadi.antarbarang/database"
	_ "com.ddabadi.antarbarang/redis"
	"com.ddabadi.antarbarang/router"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// migration()

	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%v", port)

	r := router.InitRouter()

	server := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("Server started")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}

func migration() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}
	n, err := migrate.Exec(database.GetConn(), "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Failed migration !!! ", err.Error())
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
