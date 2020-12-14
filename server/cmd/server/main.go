package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/o1111001/virtual-disk-management/server/db"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "vdm",
		User:       "vdm",
		DisableSSL: true,
		Host: 			"127.0.0.1:5432",
		Password: 	"1234",
	}

	return conn.Open()
}

func main() {
	flag.Parse()

	if server, err := ComposeApiServer(HTTPPortNumber(*httpPortNumber)); err == nil {
		go func() {
			log.Println("Starting server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize server: %s", err)
	}
}
