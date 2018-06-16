package main

import "fmt"
import "time"

func main() {

	i := 2
	//ojo, fmt.Print no mete nueva línea.
	fmt.Print("Write ", i, " as ")

	//los switch no tienen breaks
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//para casos múltiples se puede usar la coma, no en cascada como en C
	//También hay un caso default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()

	//switch sin expresión actúa como un else/if. Además las expresiones pueden ser de todo tipo, no solo literales
	switch {
	case t.Hour() < 12:
		fmt.Println("Antes de las 12")
	default:
		fmt.Println("Es por la tarde")
	}

	//existe un type switch, que compara el tipo. Se puede utilizar para conocer el tipo de un valor de interfaz.
	//en go las funciones son de primera clase (como en javascript)
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Dont know what type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
