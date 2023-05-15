package main

import (
	"fmt"

	"github.com/curso-go/variables"
)

func main() {
	estado, num := variables.ConviertoATexto(618)

	fmt.Println(estado)
	fmt.Println(num)
}
