package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sneicast/golang-crud-users-mongodb/src/services"
)

func main() {
	e := echo.New()
	e.GET("/users", services.GetUsers)
	e.GET("/users/:id", services.GetDetailUser)
	e.POST("/users", services.CreateUser)
	e.PUT("/users/:id", services.UpdatelUser)
	e.DELETE("/users/:id", services.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
