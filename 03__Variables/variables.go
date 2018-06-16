package main

import "fmt"

func main() {

	//go puede inferir el tipo automáticamente a partir del valor
	var a = "initial"
	fmt.Println(a)

	//el tipo se pone tras la declaración de la variable.
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//las variables no inicializadas se declaran con un 0 (zero-valued)
	var e bool
	var f int
	fmt.Println(e, f)

	//la sintaxis := es un sintactic sugar para declarar e inicializar a la vez una variable
	g := "short"
	fmt.Println(g)
}
