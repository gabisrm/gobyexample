package main

import "fmt"
import "math"

//las interfaces declaran una colección de cabeceras/interfaces de métodos:
// más info: http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
//Ej: interfaz para figuras geométricas:
type geometry interface {
	//solo incluyen las declaraciones, no las definiciones de los métodos
	area() float64
	perim() float64
}

//Vamos a implementar esta interfaz sobre el tipo rect y circle:
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//para implementar la interfaz sobre un tipo, solo hay que definir todos los métodos sobre ese tipo:
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//se pueden declarar variables con el tipo interfaz, que podrán hacer uso de los métodos declarados en la interfaz; sin tener por qué saber qué tipo concreto se usará
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//dado que tanto el tipo circulo y rect implementan la interfaz geometry, pueden ser argumentos de entrada de la función measure
	measure(r)
	measure(c)
}
