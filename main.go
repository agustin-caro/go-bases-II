package main

import (
	"fmt"

	"github.com/agustin-caro/go-bases-II/ejercicios/maniana"
	"github.com/agustin-caro/go-bases-II/ejercicios/tarde"
)

func main() {

	fmt.Println("hola")
	//probando paso por referencia
	number := 1
	param(&number)
	fmt.Println(number)

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio1 : ")
	maniana.Ejercicio1()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio2 : ")
	maniana.Ejercicio2()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio3 : ")
	maniana.Ejercicio3()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio4 : ")
	maniana.Ejercicio4()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio5 : ")
	maniana.Ejercicio5()

	fmt.Println("\n----------------------------------------------------------------")
	fmt.Println("Ejercicio1 (tarde) : ")
	tarde.Ejercicio1()

}

//tomando puntero como parametro
func param(n *int) {
	*n = *n + 1
}
