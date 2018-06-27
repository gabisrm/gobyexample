package main

import "fmt"

/*
* Las goroutines son una mezcla de hilos y eventos tipo nodejs => sobre una CPU, las goroutiens son similares al event loop de nodejs, en la que existe un
* scheduler interno que maneja todos esos hilos para que trabajen concurrentemente y asincronamente.
* Pero en go se puede definir el número de hilos (normalmente el mismo numero de cpus), y por cada hilo genera ese scheduler con las goroutines, sin necesidad
* de hacer spawn de procesos workers como en nodejs
 */

//esta es la función
func f(from string) {
	for i := 0; i < 50; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//aquí se llama normal
	f("direct")

	//aqui se crea una goroutine. Se ejecutará de manera concurrente:
	go f("goroutine")

	//se pueden hacer goroutines de funciones anonimas:
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//estas dos funciones anteriores se estarán ejecutando asíncronamente en gorouytines separadas.

	//esta linea espera hasta que se introduzca algo por la consola. La ponemos para ver la ejecución, ya que al ejecutarse las gorutinas asincronamente, el
	//la ejecución continua, y llegaría al final del programa sin haber terminado
	fmt.Scanln()
	fmt.Println("done")
}
