package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ayushrakesh/distribkv/db"
	"github.com/ayushrakesh/distribkv/web"
)

var (
	dbLocation = flag.String("db-location", "", "This is the url to the database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP host")
)

func parseFlags() {
	flag.Parse()
	if *dbLocation == "" {
		log.Fatalf("DB url must not be empty")
	}
}
func main() {
	parseFlags()

	db, closeFunc, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("New Database(%q): %v", *dbLocation, err)
	}

	defer closeFunc()

	server := web.NewServer(db)

	http.HandleFunc("/get", server.GetHandler)
	http.HandleFunc("/set", server.SetHandler)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))

}
