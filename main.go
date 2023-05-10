package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Pet struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    string `json:"age"`
	Breed  string `json:"breed"`
	Owner  string `json:"owner"`
	Weight string `json:"weight"`
	Color  string `json:"color"`
}

type Owner struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}

type Service struct {
	Id          int    `json:"id"`
	Name        string `json:"name"` //medical checkup,vaccine, operation
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Reservation struct {
	Id      int    `json:"id"`
	Date    string `json:"date"` //medical checkup,vaccine, operation
	Time    string `json:"time"`
	Pet     string `json:"pet"`
	Service string `json:"service"`
}

type Payment struct {
	Id             int    `json:"id"`
	Reservation    string `json:"reservation"` //medical checkup,vaccine, operation
	Price          string `json:"price"`
	Payment_Method string `json:"payment_method"`
}

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
