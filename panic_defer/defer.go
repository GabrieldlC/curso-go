package panic_defer

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("este es el primer mensaje")
	defer fmt.Println("este es el último mensaje")

	fmt.Println("este es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrió un error que genderó panic\n%v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("el valor de a es 1")
	}
}
