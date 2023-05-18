package main

import (
	"fmt"

	"github.com/curso-go/goroutines"
)

func main() {
	canal1 := make(chan bool)

	go goroutines.MiNombreLentooo("Gabriel de la Cuadra", canal1)

	fmt.Println("A ver si escribís más rápido tu nombre")
	var x string
	fmt.Scanln(&x)

	<-canal1
}
