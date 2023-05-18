package arreglos_slice

import "fmt"

var tabla [10]int = [10]int{10, 0, 55}
var matriz [20][30]int

func MuestroArray() {
	tabla[7] = 77
	tabla[2] = 22

	tabla2 := [10]string{"gabriel", "juan"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}
