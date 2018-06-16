package main

import "fmt"

func main() {

	//cuando es una unica condicion, no se pone paréntesis en la premisa del if. Siempre se pone los brackets, y no existe un operador ternario como en C
	if 7%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}

	//se puede preceder la premisa por declaraciones. Las variables declaradas en la premisa están disponibles a través de todo el bloque
	if num := 9; num < 0 {
		fmt.Println("Es negativo")
	} else if num < 10 {
		fmt.Println("Tiene  un dígito")
	} else {
		fmt.Println("Tiene más de un")
	}

}
