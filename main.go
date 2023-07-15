package main

import (
	"fmt"
)

import "github.com/labstack/echo/v4"

// IT IS How a Class
type Car struct {
	Name string
	Price int
}

// It is How a Method 
func (c Car) Andar() {
	fmt.Println("ESTAMOS ANDANDO", c.Name, "agora")
}

var cars []Car

func createCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Porsche", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 100})
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func getTest(c echo.Context) error {
	return c.String(200, "Estamos online")
}

func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.GET("/", getTest)
	createCars()
	fmt.Println(cars, e)
	e.Logger.Fatal(e.Start(":3001"))
	  // http.Handle(*"estamos online", nil)
		// Não deu certo
}
// The Go is typing strongly language
// For this is need to type params and return of function
func soma(param1, param2 int) (int, error) {
// This format typings é the following, int is the first return the function and error is the second return
	if param1 > 10 {
		return 0, fmt.Errorf("Número muito alto")
	}
	return param1 + param2, nil
}