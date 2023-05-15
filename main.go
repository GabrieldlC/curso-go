package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, num := variables.ConviertoATexto(618)

	fmt.Println(estado)
	fmt.Println(num)*/

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
