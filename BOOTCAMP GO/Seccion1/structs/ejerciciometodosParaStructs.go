package main

import "fmt"

//Los métodos son funciones que se ejecutan vía una variable del tipo de un dato.
//El identificador sirve para mandar a llamar cualquier función con la estructura.
type Usuarios struct{ 
	edad int
	nombre, apellido string 
}

/

func (usuario Usuarios) nombre_completo() string { //estructura de nuestra función.
	//usuario= identificador de la esctructura para utilizarlo dentro de la función.
	//Usuarios= estructura a la cual le estamos agregando el método.
	//nombre_completo()= nombre del método.
	//string= tipo de datos que va a retornar.
	return usuario.nombre + " " + usuario.apellido

}




func main() {
	silvia := new(Usuarios)

	silvia.nombre = "Silvia"
	silvia.apellido = "González"
	

	fmt.Println(silvia.nombre_completo())
}
