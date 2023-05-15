package ejercicios

import "strconv"

func MayorMejorQueCien(txt string) (int, string) {
	num, err := strconv.Atoi(txt)

	if err != nil {
		return 0, "hubo un error: " + err.Error()
	}
	if num < 100 {
		return num, "Es menor a 100"
	}
	return num, "Es mayor a 100"
}
