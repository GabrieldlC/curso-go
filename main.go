package main

import (
	"github.com/curso-go/middleware"
)

func main() {
	// canal1 := make(chan bool)

	// go goroutines.MiNombreLentooo("El Valen", canal1)

	// fmt.Println("A ver si escribís más rápido tu nombre")
	// var x string
	// fmt.Scanln(&x)

	// <-canal1

	// webserver.MiWebServer()

	middleware.MiMiddleware()
}
