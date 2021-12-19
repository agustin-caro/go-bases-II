package tarde

type Alumno struct {
	nombre   string
	apellido string
	dni      string
	fecha    Fecha
}

type Fecha struct {
	dia, mes, anio int
}

func Ejercicio1() {

	var a1 Alumno
	a1.nombre = "Agustin"
	a1.apellido = "Caro"

	f1 := Fecha{10, 10, 1994}

	a1.fecha = f1

}
