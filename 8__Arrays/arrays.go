package main

import "fmt"

/*
* NOTA: la gran diferencia entre arrays en Go y en C es que en Go el array es un valor en sí! Si se pasa el array por argumento de entrada, se pasa
* una copia del array, no un puntero al primer elemento del array!. Es como una estructura de datos indexada.
* En Go se utiliza más los slices, que son mas flexibles (siguiente ejemplo).
 */

func main() {

	//el tipo es [5]int, es decir que la longitud del array pertenece a la descripción del tipo. Por defecto se inicializan a 0
	var a [5]int
	//NOTA: fmt.Println actúa como console.log(), en el sentido de que es una función variádica
	fmt.Println("emp:", a)

	//asignación como toda la vida. indices empiezan en 0
	a[4] = 100
	fmt.Println("set:", a)

	//obtención de valores como toda la vida
	fmt.Println("get:", a[4])

	//función len incluida en el builtin del sistema (donde se incluyen funciones como new(), panic(), la declaración de los tipos.... no hace falta importar)
	fmt.Println("len:", len(a))

	//declaración e inicialización de un array en una línea
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dlc:", b)

	//arrays multidimensionales
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D: ", twoD)

}
