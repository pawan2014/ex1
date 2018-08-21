package handler

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestStoreSuite(t *testing.T) {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	InitStore(&dbStore{mydb: db})
	v, err := store.GetMachines()
	if err != nil {
		panic(err)
	}
	fmt.Print(v)

}
