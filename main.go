package main

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:true`
	Message string      `json:message`
	Data    interface{} `json:data`
}

func main() {
	e := echo.New()
	e.GET("/pets", GetPetsController)
	e.Start(":8000")
}
func GetPetsController(c echo.Context) error {
	return c.JSON(200, BaseResponse{
		true, "success", "hello",
	})

}
