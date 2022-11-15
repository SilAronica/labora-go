package main

import "fmt"
	//Las estructuras son mutables: podemos modifica los valores de los campos después de haber instanciado.
	//Type = indica que va a definir su nuevo tipo
	//User = nomre del tipo
	//struct = indica que vamos a definir una estructura.

type User struct{ //Las llaves permiten enlistar los atributos o propiedades que va a tener la estructura, seguido por el tipo de dato, por ejemplo:
	edad int
	nombre, apellido string //estos campos se pueden declarar por separado, pero tb en una sola línea porque ambos son el mismo tipo.
}

//func main() {
	//silvia := User{nombre: "Silvia", apellido: "Gonzalez", edad: 38}
	
	//fmt.Println(silvia)
//}

//En Go cuando queremos acceder a los campos de una estructura a través de un puntero, tendríamos que extraer el dato del puntero utilizando el asterisco, punto y un campo. 
func main() {
	silvia := new(User)
	//La función new nos permite declarar una estructura pero en este caso esta función retorna a un puntero

	fmt.Println(silvia)
}
