package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type BaseResponse struct {
	Status  bool        `json:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	// Connection parameters
	connStr := "user=postgres password=admin123 dbname=db_prakerja sslmode=disable"

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	// Perform database operations...
	e := echo.New()
	e.GET("/pets", GetPetsController)
	e.Start(":8000")
}
func GetPetsController(c echo.Context) error {
	return c.JSON(200, BaseResponse{
		true, "success", "hello",
	})

}
