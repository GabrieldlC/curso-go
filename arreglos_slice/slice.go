package arreglos_slice

import "fmt"

var arr []int = []int{2, 5, 4}
var vector [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestroSlice() {
	fmt.Println(arr)

	porcion := vector[3:]   //Slice creado desde la posición 3 de un vector
	porcion2 := vector[:5]  //Slice creado desde pos 0 hasta 5 de un vector
	porcion3 := vector[2:7] //Slice creado desde pos 2 hasta 7 de un vector

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
