package main

import "fmt"

//go sporta punteros y referencias
//igual que en C. PERO OJO, no es como C++ => no existe paso por referencia!! por lo que no existen variables referencias.
//Para hacer paso por referencia, siempre punteros.
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 5
}

func valuePass(iptr *int) {
	iptr = nil
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	//nada, le pasamos una copia; no modifica el original
	zeroval(i)
	fmt.Println("zeroval:", i)

	//le pasamos la dirección de memoria de i => lo modifica. Basicamente sigue siendo una llamada por valor, a la cual se le pasa una copia de la posición de memoria
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//hace una copia de valor del puntero, por lo que aunque se ponga a nulo en la función, se pone a nulo ese puntero, no a i.
	valuePass(&i)
	fmt.Println("valuePass:", i)

}
