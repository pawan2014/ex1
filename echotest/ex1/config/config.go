package config

import "database/sql"

type DbConfig struct {
	mydb *sql.DB
}

func NewConfig(url string) DbConfig {
	db, err := sql.Open("postgres", url)
	CheckErr(err)

	return DbConfig{db}
}
func (c DbConfig) Close() {
	c.mydb.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Hello() {

}
