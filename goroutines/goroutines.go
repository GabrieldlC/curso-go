package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Print(letra)
	}
	canal1 <- true
}
