package main

import "fmt"

func main() {
	matriz := []int{1,2}
		if matriz == nil{
			fmt.Println("Esta vacío")
	}else {
		fmt.Println(matriz)
	}
}



//Su base es un arreglo, pero tiene otras capacidades que lo hacen más dinámico.
//Puntero al arreglo
//Longitud del arreglo al que apunta.
//Capacidad.
//Nil = es mucho menos común. Hay un valor cero para todos los datos. Valor Nulo como valor.