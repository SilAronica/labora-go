//Cuando recibimos la estructura en sí como un puntero en lugar de como una copia. 
//Recordar que cada vez que pasamos un argumento a una función en GO, este argumento se pasa como una copia: duplica el objeto y lo envía. 
//Pasar la referencia a la estructura como puntero es una operación más económica, no es lo mismo copiar una estructura que copiar una referencia a memoria.
//se puede agregar métodos solamente a las estructuras declaradas en el mismo paquete, si se estuviera importando la etructura desde otro paquete, no podríamos agregarle otras funcionalidades.

package main

import "fmt"

type Usuarios struct{ 
	edad int
	nombre, apellido string 
}



func (this Usuarios) nombre_completo() string { 
	return this.nombre + " " + this.apellido

}

func (this *Usuarios) set_name(n string) {
	this.nombre = n

}




func main() {
	silvia := new(Usuarios)

	silvia.nombre = "Silvia"
	
	silvia.set_name("Verónica")

	fmt.Println(silvia.nombre)
}
