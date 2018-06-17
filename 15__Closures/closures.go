package main

import "fmt"

/*
* Go soporta la declaración de funciones anónimas, como en js, que también crean su propio closure
 */

//función que devuelve una función que devuelve un int
func intSeq() func() int {
	i := 0

	//función anónima en el cuerpo de intSeq. Dicha función hace un closure con i, es decir, que la captura
	return func() int {
		i++
		return i
	}
}

func main() {

	//asignamos la función a la variable nextInt. Cada vez que se llame, actualizará el valor de i capturado
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//aquí volvemos a crear otro closure, independiente al anterior
	newInts := intSeq()
	fmt.Println(newInts())
}
