package main

import (
	"fmt"
)

/*
* los channels son los pipes para interconectar datos entre goroutines.fmt* Se pueden insertar datos en un channel desde una goroutine,
* y recibirlos en otra goroutine
 */

func main() {
	//un canal se crea con make(chan <T>), siendo T el tipo de valores que pasa el canal.
	messages := make(chan string)

	//para pasar un valor por un canal usamos la sintaxis: canal <- valor. En este caso enviamos "ping"
	go func() { messages <- "ping" }()

	//fmt.Scanln()

	//asi recibimos valores de un canal
	msg := <-messages
	fmt.Println(msg)

	//por defecto, los envios y recepciones se bloquean hasta que ambas partes estén preparadas (declaradas y llegados a ese punto de ejecución), asi se
	//evita la necesidad de sincronizar (probar a descomentar fmt.Scanln())
}
