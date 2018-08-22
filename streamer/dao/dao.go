package dao

import (
	"database/sql"
	"log"

	"github.com/ex1/streamer/model"
)

// Dao type for all differetn dao like mySql, Local
type Dao interface {
	Hi() string
	GetMachinesAndTags() []model.Machine
	GetMachine(id string) model.Machine
}

// PgDao for attaching all the Dao methods applicable for postgres
type PgDao struct {
	mydb *sql.DB
}

func (p PgDao) Hi() string {
	return "hello"
}
func (p PgDao) GetMachinesAndTags() []model.Machine {
	var mydata []model.Machine
	rows, err := p.mydb.Query("SELECT machineid, machinename FROM machine")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		mach := model.Machine{}
		err := rows.Scan(&mach.Machineid, &mach.MachineName)
		if err != nil {
			log.Fatal(err)
		}
		myTagRows, _ := p.mydb.Query(`SELECT tagid, tagtype, tagname, formula, 
			frequency FROM public.machinetags  where machineid=$1`, mach.Machineid)
		var tagesTempData []model.Tag
		for myTagRows.Next() {
			tag := model.Tag{}
			err := myTagRows.Scan(&tag.TagID, &tag.TagType, &tag.TagName, &tag.Formula, &tag.Frequency)
			if err != nil {
				log.Fatal(err)
			}
			tagesTempData = append(tagesTempData, tag)
		}

		mach.Tags = tagesTempData
		mydata = append(mydata, mach)
	}
	return mydata
}
func (p PgDao) GetMachine(id string) model.Machine {
	var mydata model.Machine
	myTagRows, _ := p.mydb.Query(`SELECT tagid, tagtype, tagname, formula, 
			frequency FROM public.machinetags  where machineid=$1`, id)
	var tagesTempData []model.Tag
	for myTagRows.Next() {
		tag := model.Tag{}
		err := myTagRows.Scan(&tag.TagID, &tag.TagType, &tag.TagName, &tag.Formula, &tag.Frequency)
		if err != nil {
			log.Fatal(err)
		}
		tagesTempData = append(tagesTempData, tag)
	}

	mydata.Tags = tagesTempData
	mydata.Machineid = id

	return mydata
}

// NewPgDao for postgress sql
func NewPgDao(db *sql.DB) Dao {
	return &PgDao{db}
}

//
