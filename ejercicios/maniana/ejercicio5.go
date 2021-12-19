package maniana

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Ejercicio5() {

	fmt.Printf("5 perros: ")
	funcPerro, error := Animal(perro)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("%v gramos de alimento", funcPerro(5))

	fmt.Printf("\n1 hamster: ")
	funcHamster, error := Animal(hamster)
	if error != nil {
		fmt.Println(error)
	} else {

		fmt.Printf("%v gramos de alimento", funcHamster(1))
	}

	fmt.Printf("\nJirafa (no existe) :  ")
	funcJirafa, error := Animal("jirafa")
	if error != nil {
		fmt.Println(error)
	} else {

		fmt.Printf("%v gramos de alimento", funcJirafa(1))
	}

}

func Animal(tipo string) (func(cantidad int) int, error) {

	switch tipo {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("animal inexistente")
	}
}

func animalPerro(cantidad int) int {
	return cantidad * 10000
}

func animalGato(cantidad int) int {
	return cantidad * 5000
}

func animalHamster(cantidad int) int {
	return cantidad * 250
}

func animalTarantula(cantidad int) int {
	return cantidad * 150
}
