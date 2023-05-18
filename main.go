package main

import (
	e "github.com/curso-go/ejer_interfaces"
	"github.com/curso-go/modelos"
)

func main() {
	elMati := new(modelos.Hombre)
	e.HumanosRespirando(elMati)

	laMaria := new(modelos.Mujer)
	e.HumanosRespirando(laMaria)

	elPepe := new(modelos.Mujer)
	e.SerVivoEstaVivo(elPepe)
}
