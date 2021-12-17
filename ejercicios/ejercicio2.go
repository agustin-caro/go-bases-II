package ejercicios

import (
	"errors"
	"fmt"
)

func Ejercicio2() {

	promedio, error := promedio(-1, 2, 3)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Promedio: ", promedio)
	}
}

func promedio(numbers ...int) (float64, error) {

	var promedio float64
	for _, n := range numbers {
		if n < 0 {
			return 0, errors.New("Numero negativo")
		}

		promedio += float64(n)
	}
	return promedio / float64(len(numbers)), nil
}
