package handler

import (
	"database/sql"

	"github.com/ex1/echotest/ex1/model"
)

type dbStore struct {
	mydb *sql.DB
}

type MyStore interface {
	GetMachines() ([]*model.Machine, error)
}

func (store *dbStore) GetMachines() ([]*model.Machine, error) {
	machines := []*model.Machine{}
	rows, err := store.mydb.Query("SELECT * FROM machine")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var machineid string
		var machinename string
		err = rows.Scan(&machineid, &machinename)
		if err != nil {
			panic(err)
		}
		machines = append(machines, &model.Machine{Machineid: machineid, MachineName: machinename})
	}

	return machines, nil
}

var store MyStore

func InitStore(s MyStore) {
	store = s
}
