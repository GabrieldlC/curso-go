package ejer_interfaces

import (
	"fmt"

	"github.com/curso-go/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.Sexo())
}

func SerVivoEstaVivo(sv interfaces.SerVivo) {
	fmt.Printf("Soy un ser vivo o muerto? %t \n", sv.EstaVivo())
}
