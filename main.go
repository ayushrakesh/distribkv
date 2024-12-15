package main

import (
	"flag"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var (
	dbLocation = flag.String("db-location", "", "This is the url to the database")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("DB url must not be empty")
	}
}
func main() {
	parseFlags()

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(db.Info())
}
