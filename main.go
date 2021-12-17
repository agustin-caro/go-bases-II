package main

import (
	"fmt"

	"github.com/agustin-caro/go-bases-II/ejercicios"
)

func main() {

	fmt.Println("hola")
	//probando paso por referencia
	number := 1
	param(&number)
	fmt.Println(number)

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio1 : ")
	ejercicios.Ejercicio1()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio2 : ")
	ejercicios.Ejercicio2()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio3 : ")
	ejercicios.Ejercicio3()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio4 : ")
	ejercicios.Ejercicio4()
}

//tomando puntero como parametro
func param(n *int) {
	*n = *n + 1
}
