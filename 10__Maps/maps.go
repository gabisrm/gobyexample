package main

import (
	"fmt"
)

/*
* Los maps son como los objetos asociativos de javascript, dicts de python o hashes de ruby
 */

func main() {

	//para declarar un map => make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 10

	fmt.Println("map:", m)

	//obtener un valor a partir del key
	v1 := m["k1"]
	fmt.Println("v1", v1)

	//len también funciona con los maps
	fmt.Println("len", len(m))

	//para eliminar pares clave/valor
	delete(m, "k2")
	fmt.Println("map:", m)

	//en realidad el get devuelve dos valores, primero el valor, y el segundo valor es un booleano que indica si se ha encontrado dicha clave o no. Se usa para saber si de verdad no existe o que existe pero el valor es nulo (0, "")
	//Para indicar que no se quiere un valor, se usa el 'blank identifier' => _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//aqui blank devuelve 0, por eso se utiliza el segundo valor devuelto para desambiguar si es en realidad un valor válido o no
	blank := m["k2"]
	fmt.Println("blank:", blank)

	//declarar e inicializar en una línea
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
