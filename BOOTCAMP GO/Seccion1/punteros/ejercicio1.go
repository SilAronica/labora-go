package main 

import "fmt"

func main() {
	//El puntero es una dirección de memoria.
	//En lugar de obtener el valor obtenemos la dirección en la que está el valor.
	//T => al tipo de dato: *int, *string, *Struct.
	//Valor zero = nil

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)
}