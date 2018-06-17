package main

import "fmt"

//función que toma dos int como argumento de entrada y devuelve otro int
//Recordar primero variable y después tipo
func plus(a int, b int) int {
	//necesita un return explicito, no es como ruby
	return a + b
}

//se puede omitir el tipo de las primeras si tenemos argumentos consecutivos del mismo tipo
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
