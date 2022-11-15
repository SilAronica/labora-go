package main

import "fmt"

func main() {
	
	var a, b int
	
	//Entrada
	fmt.Print("Ingrese número 01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese número 02: ")
	fmt.Scanln(&b)

	//Igualdad ==
	fmt.Println("¿Son Iguales? =>", a == b)

	//Distintos !=
	fmt.Println("¿Son Distintos? =>", a != b)

	//Menor que <
	fmt.Println("El Primero es menor que el segundo =>", a < b)

	//Menor o igual <=
	fmt.Println("El Primero es menor o igual que el segundo =>", a <= b)

	//Mayor que >
	fmt.Println("El Primero es mayor que el segundo =>", a > b)

	//Mayor o igual >=
	fmt.Println("El Primero es mayor o igual que el segundo =>", a >= b)
	

}