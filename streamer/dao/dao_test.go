package dao

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestGello(t *testing.T) {
	//t.Error("Expected 1.5, got ", 1)
	// t.Error(
	// 	"For", "Pawan",
	// 	"expected", "Kumar",
	// 	"got", "None",
	// )
}

func TestGetAllMachinesTags(t *testing.T) {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err1 := db.Ping()
	if err1 != nil {
		//do something here
	}
	mydao := NewPgDao(db)
	machines := mydao.GetMachinesAndTags()
	for _, v := range machines {
		fmt.Println(v.MachineName)
		for _, k := range v.Tags {
			fmt.Println("\t", k.TagID)
		}
	}
}
func TestGetMachine(t *testing.T) {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err1 := db.Ping()
	if err1 != nil {
		//do something here
	}
	mydao := NewPgDao(db)
	machine := mydao.GetMachine("1")

	fmt.Println(machine.Machineid)
	for _, k := range machine.Tags {
		fmt.Println("\t", k.TagID)
	}

}
