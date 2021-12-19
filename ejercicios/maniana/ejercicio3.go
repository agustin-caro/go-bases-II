package maniana

import "fmt"

func Ejercicio3() {

	minutos := 60
	categoria := "C"

	fmt.Printf("Salario: %.1f", calcularSalario(minutos, categoria))

}

func calcularSalario(minutos int, categoria string) float64 {
	var salario float64
	salario = 0
	horas := float64(minutos / 60)
	switch {
	case categoria == "A":
		salario += horas * 3000 * 1.5
	case categoria == "B":
		salario += horas * 1500 * 1.2
	case categoria == "C":
		salario += horas * 1000

	}

	return salario
}
