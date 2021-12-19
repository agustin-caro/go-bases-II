package maniana

import (
	"errors"
	"fmt"
)

const (
	minimo    = "minimo"
	promedio2 = "promedio"
	maximo    = "maximo"
)

func Ejercicio4() {

	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio2)
	maxFunc, err := operacion(maximo)

	if err != nil {
		fmt.Printf("Error")
	}

	valorMinimo, _ := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio, _ := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Valores min prom y max : %.1f - %.1f - %.1f", valorMinimo, valorMaximo, valorPromedio)

}

func operacion(operacion string) (func(numbers ...int) (float64, error), error) {

	switch operacion {
	case minimo:
		return minFunc, nil
	case promedio2:
		return promFunc, nil
	case maximo:
		return maxFunc, nil
	default:
		return nil, errors.New("invalid operacion")
	}
}

func minFunc(numbers ...int) (float64, error) {
	min, _ := findMinAndMax(numbers...)

	return float64(min), nil
}

func maxFunc(numbers ...int) (float64, error) {
	_, max := findMinAndMax(numbers...)

	return float64(max), nil
}

func promFunc(numbers ...int) (float64, error) {
	return promedio(numbers...)
}

func findMinAndMax(numbers ...int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
