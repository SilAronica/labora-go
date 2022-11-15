//Gorutines son como los hilos pero más ligeras. Las funciones cuando están en gorutines separadas van a estar ejecutándose, mientras el hilo principal sigue con sus instrucciones. Conceptualmente, porque no siempre van a ejecutarse en paralelo. No tenemos que esperar que otra función termine para poder seguir con el trabajo
//Cómo se comunican las gorutines?, con los canales. Una gorutine manda un objeto y hay otras que escucan ese canal y obtienen de ahí lo que necesitan.package concurrencias

//Cómo es la coodinación? como select: lo que permite es estar escuchando diferentes canales y ver cuál está listo y entonces tomar uno u otro para hacer lo que necesita.

package main

import (
	"fmt"
	"os"
)

func main(){
	f, err := os.Create("mi_archivo.txt")//Creo el archivo.
	if err != nil{//Chequeo errores.
		panic(err)//Si existe erro lanzo un panic.
	}

	final := 16777215
	for i := 0; i <= final; i++ {//Itero sobre todos los numeros hasta llegar a 16777215.
		_, err = f.WriteString(fmt.Sprintf("%60x\n", i))//Escibo su versión en he
		if err != nil{
			panic(err)
		}
	}
}
