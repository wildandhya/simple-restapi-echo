package main

import (
	"go-echo/db"
	"go-echo/helper"
	"go-echo/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.HTTPErrorHandler = helper.ErrHandler
	e.Logger.Fatal(e.Start(":9020"))
}
