package main

import (
	"net/http"

	"github.com/ex1/echotest/ex1/config"
	"github.com/ex1/echotest/ex1/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {

	//
	config.Hello()
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	myDBConfig := config.NewConfig(connStr)
	// var mydb *sql.DB
	// mydb, err := sql.Open("postgres", connStr)
	// CheckErr(err)
	defer myDBConfig.Close()

	//
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// create a handler
	h := &handler.Handler{}

	e.GET("/machines", h.GetMachines)
	e.GET("/machine/:id", h.GetMachineWithTags)
	e.GET("/machine/:id/sstatus", h.MachineStreamingStatus)
	// Add & Update  Machine
	e.POST("/machine", h.AddMachineWithTags)
	e.PUT("/machine", h.AddMachineWithTags)

	// Add & Update Tag for a machine
	e.POST("/machine/:id/tag", h.UpdateTagUnderMachine)
	e.PUT("/machine/:id/tag/:tagid", h.UpdateTagUnderMachine)
	// Stream Status
	e.POST("/machine/:id/start", h.StartMachineStream)
	e.POST("/machine/:id/stop", h.StopMachineStream)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
