package funciones

import "fmt"

func tabla(valor int) func() int {
	secuencia := 0
	return func() int {
		secuencia++
		return valor * secuencia
	}
}

func LlamarClosure() {
	tablaDel := 2
	MiTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
