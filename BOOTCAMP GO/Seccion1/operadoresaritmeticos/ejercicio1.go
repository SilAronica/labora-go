package main

import "fmt"

func main() {
	var a, b int

	//Entrada
	fmt.Print("Ingrese número 01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese número 02: ")
	fmt.Scanln(&b)

	//Proceso
	suma := a + b
	resta := a - b 
	multi := a * b
	divi := a / b 
	mod := a % b 

	//Salida
	fmt.Println("La Suma es: ", suma)
	fmt.Println("La Resta es: ", resta)
	fmt.Println("La Multiplicación es: ", multi)
	fmt.Println("La División es: ", divi)
	fmt.Println("El Porcentaje es: ", mod)

}