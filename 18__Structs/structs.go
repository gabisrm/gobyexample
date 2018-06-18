package main

import "fmt"

//los structs se definen de la siguiente manera:
type person struct {
	name string
	age  int
}

func main() {

	//y se declaran de la siguiente manera:
	//rápido:
	fmt.Println(person{"Bob:", 20})

	//identificando los campos
	fmt.Println(person{name: "Alice", age: 20})

	//los campos que no se declaren se inicializan a 0
	fmt.Println(person{name: "Fred"})

	//puntero a la estructura
	fmt.Println(&person{name: "ann", age: 89})

	s := person{name: "Sean", age: 90}
	//se acceden con el operador punto
	fmt.Println(s.name)

	sp := &s
	//los punteros a structs se derreferencian automáticamente:
	fmt.Println(sp.age)

	//son mutables; se pueden modificar sus valores tras la creación
	sp.age = 53
	fmt.Println(sp.age)

}
