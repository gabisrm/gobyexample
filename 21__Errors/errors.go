package main

/*
* En go, los errores se comunican a través del segundo valor de return; no existe try/catch block! Si que existe un tipo 'Error', que es
* el tipo de datos que se devuelve por ese segundo return.
* https://blog.golang.org/error-handling-and-go
 */

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		//dado que se sigue teniendo que devolver dos valores, el primer valor lo ponemos como -1, o con lo que se quiera.
		//el segundo valor tienen que implementar la interfaz tipo Error
		return -1, errors.New("can't work with 42")
	}
	//en caso de no error, segundo return a nil
	return arg + 3, nil
}

//se pueden crear tipos customizados de error, SIEMPRE que implementen el método Error(), ya que es el requerido por el interface Error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it!"}
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {

		//en el propio if se puede asignar y hacer un inline check. Es como se chequea normalmente en Go en una misma línea
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//fuerzo el error
	_, e := f2(42)

	//obtengo los datos del error pasándolo a una instancia del struct que me he creado -> TYPE ASSERTION
	//type assertion da acceso a los valores correspondientes a la interfaz de una instancia en concreto
	//Ej: t := i.(T)  -> esto hace un assert de que el valor de la interfaz de i es de tipo T, y asigna las propiedades de T a t; en caso de que no -> PANIC
	//como segundo valor de return devuelve un booleano indicando si el assertion ha succeded -> t, ok := i.(T)
	//https://tour.golang.org/methods/15
	if ae, ok := e.(*argError); ok {
		//en caso de que e implementa la interfaz argError, obtengo los datos
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
