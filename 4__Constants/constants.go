package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000000000

	const d = 3e20 / n
	//una constante numérica no tiene tipo hasta que se le da uno (algo así como los defines en c)
	fmt.Println(d)
	fmt.Println(int64(d))

	//por ehemplo aquí se le daría el tipo float64, porque es el tipo del argumento de entrada de math.Sin
	fmt.Println(math.Sin(n))
}
