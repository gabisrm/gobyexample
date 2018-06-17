package main

import "fmt"

//los múltiples valores de retorno se usan mucho para devolver el valor y el error (en caso de que haya)

//para declarar múltiples valores de retorno, se ponen los tipos entre paréntesis
func vals() (int, int) {

	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	//si solo se quiere un subset, se utiliza el blank identifier -> _
	_, c := vals()
	fmt.Println(c)

}
