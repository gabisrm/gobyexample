package main

import "fmt"

//en Go no se definen clases explicitamente, sino tipos y métodos asociados a los tipos.
//definimos primero el tipo
type rect struct {
	width, height int
}

//el siguiente método tienen un 'receiver type' de un *rect (el receiver type se pone antes del nombre de la función):
func (r *rect) area() int {
	return r.width * r.height
}

//los métodos se pueden definir para receptores tipo punteros o valores
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	//go maneja automáticamente la conversión entre punteros y valores en la llamada de los métodos:
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	//permite trabajar indiferente entre punteros y valores, quedando la elección solamente por cuestiones de funcionalidad(si quiero mutar el objeto o si no quiero copiar el valor tan grande en memoria uso punteros)
}
