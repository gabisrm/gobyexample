package main

import (
	"fmt"
)

//ver: https://blog.golang.org/go-slices-usage-and-internals:
/*
* En realidad un slice es un descriptor de un segmento de un array, en la que se almacena el puntero al array, la longitud y la capacidad.
* La capacidad se puede especificar también con make() => func make([]T, len, cap) []T  , donde T es el tipo
*
*
 */
//slices son un tipo literal de go, que proporciona una interfaz más potente que los arrays y funcionalidades de secuenciación
func main() {

	//los slices se declaran indicando el tipo de datos que van a albergar y no los datos en sí. Por defecto son de longitud 0. Si se quiere especificar la longitud, se usa el builtin make()
	//se inicializa a 0
	s := make([]string, 3) //cuando no se indica la capacidad, se toma la misma que la longitud
	fmt.Println("emp:", s)

	//podemos manipularlos como si fueran arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	//soporta otras operaciones como append(). Es como si fuera una subclase de un array con más operaciones
	//append devuelve un nuevo slice con lo añadido, pero no modifica el original.
	//OJO: no se puede aumentar la capacidad de un slice. Para incrementar un slice, se debe crear uno nuevo!!
	//append incrementará el tamaño del slice si hace falta. Pero si no hace falta no. Ej: sse ha declarado un slice con capacidad 10, de las cuales se han rellenado 5. Append ahí no incrementará a no ser que se agreguen más de 5 más
	s = append(s, "d", "e", "f")
	t := append(s, "g")

	fmt.Println("s:", s)
	fmt.Println("t:", t)

	//copiar slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//operador 'slice' => slice[low:high). es como en python
	l := s[2:5]
	fmt.Println("sl1:", l)

	//hasta 5 pero sin incluir:
	l = s[:5]
	fmt.Println("sl2:", l)

	//desde el 2
	l = s[2:]
	fmt.Println("sl3:", l)

	/*
	* NOTA: Slicing no copia los datos del slice, sino que inicializa un nuevo slice donde el puntero al array de datos se copia, por lo que apunta al mismo
	* contenido del original => modificar en la copia modifica en el original!!
	 */
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	//declaración e inicialización en una línea
	u := []string{"g", "h", "i"}
	fmt.Println("dlc:", u)

	//slices multidimensionales. La longitud de las otras dimensioens pueden variar (como python). Esto no se puede hacer con arrays normales
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	/*
	* Con slices hay que tener en cuenta el dejarlos en memoria => el recolector de basura solo pasa cuando no tiene ninguna referencia. OJO con
	* devolver de una función una referencia a una parte de un slice mayor, ya que en memoria se quedará el slice entero, aunque se use solo una pequeña parte.
	* La solución a esto es devolver una copia de esa parte pequeña con copy() (ver la seccuión de gotchas en la url de arriba)
	 */
}
