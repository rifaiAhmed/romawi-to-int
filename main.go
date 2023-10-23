package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi server Echo
	e := echo.New()

	// Rute sederhana untuk menangani permintaan
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Selamat datang di server Echo Golang!"})
	})
	e.GET("/:romawi", func(c echo.Context) error {
		romawi := strings.ToUpper(c.Param("romawi"))
		convert := RomanToInteger(romawi)

		response := map[string]string{"message": romawi + " = " + fmt.Sprint(convert)}

		return c.JSON(http.StatusOK, response)
	})

	// Mulai server di port 8080
	e.Start(":8080")
}

func RomanToInteger(roman string) int {
	romanValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	result := 0

	for i := 0; i < len(roman); i++ {
		currentValue := romanValues[roman[i]]
		if i+1 < len(roman) && romanValues[roman[i+1]] > currentValue {
			result -= currentValue
		} else {
			result += currentValue
		}
	}

	return result
}
