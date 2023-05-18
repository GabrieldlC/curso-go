package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeNum() {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error
	fmt.Println("Ingrese su número")

	for {
		scanner.Scan()
		num, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Ingrese un número válido")
		} else {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(num * i)
	}
}
