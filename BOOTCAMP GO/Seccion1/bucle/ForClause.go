//for [ Initial Statement ] ; [ Condition ] ; [ Post Statement ] {
    //[Action]
//}
//Un bucle ForClouse se define como un bucle que tiene una instrucción inicial seguida de una condición y luego una posinstrucción.package bucle
package main

import "fmt"

func main() {
	//Estructura del for 1)-for: palabra clave; 2)-Clásula for. a)-Declaración de iniciliación "init"= i, es donde vamos a inicializar por primera vez una variable; b)-Condición: nos va a decir hasta cuand voy a iterar, mientras esa condición sea verdadera voy a iterar mi código; c)-Declaración de post: declaración que se ejecuta al final de cada iteración. 3)- luego de la llave viene el bloque de código: es lo que se ejecuta en cada iteración.
	//Todas las declaraciones se separan con ;.
	for i := 0; i < 5; i++ {
		fmt.Println(i)//Se imprime el valor de i.

}

}