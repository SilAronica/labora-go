//Tenemos un ciclo externo, qe va a iterar 3 veces.package bucle
//Y luego teneos un ciclo interno, que itera 5 veces.
//1° iteación el ciclo externo, 5 iteraciones del ciclo interno.Por lo tanto, por cada iteración del externo, tenemos 5 iteraciones del interno.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {//Ciclo externo
		fmt.Printf("El ciclo externo %d\n", i )
		for j := 0; j < 3; j++ {//Ciclo interno
		fmt.Printf("\t\t El ciclo interno %d\n", j)
		}
	}
}