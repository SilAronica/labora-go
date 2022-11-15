package main

import "fmt"

func main() {

	//c := 0

	//Bucle con una sola condición
	//for c <= 5 {
		//fmt.Println(c)
		//c++
	//}


	//Bucle con más de una condición
	for i := 0; i <= 10; i++ {
		if i == 5 {
			//fmt.Println("Se rompe el bucle")
			//break
			fmt.Println("Estamos en la 5ta ejecución")
			continue
			fmt.Println("No me ejecutaré")

		}
		fmt.Println(i)
	}		
}