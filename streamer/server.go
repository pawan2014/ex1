package main

import (
	"database/sql"
	"net/http"

	"github.com/ex1/streamer/config"
	"github.com/ex1/streamer/dao"
	"github.com/ex1/streamer/handler"
	"github.com/ex1/streamer/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

//var Myservice *service.Service
var pawan string

func main() {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	dao := dao.NewPgDao(db)
	service.NewService(dao)

	// MQ Connection
	rabbitconn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer rabbitconn.Close()
	if err != nil {
		panic(err)
	}
	config.NewRabbitConfig(rabbitconn).Configure()

	// Server
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

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
