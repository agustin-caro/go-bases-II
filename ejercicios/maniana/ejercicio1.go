package maniana

import "fmt"

func Ejercicio1() {

	fmt.Println("Salario con impuestos: ")
	var salario float64
	salario = 60000

	fmt.Println(calcularDescuento(salario))

}

func calcularDescuento(salario float64) float64 {

	if salario > 50000 {
		salario -= (salario * 17) / 100
	}
	if salario > 150000 {
		salario -= (salario * 10) / 100
	}

	return salario
}
