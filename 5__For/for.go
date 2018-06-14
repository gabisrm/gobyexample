package main

import "fmt"

func main() {
	i := 1

	//bucle for con una condición, con la variable declarada fuera
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//loop clásico como en C
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//un for sin condiciones es como un for(;;) o while(1). Necesita un break o un return
	for {
		fmt.Println("loop")
		break
	}

	//tambien se puede usar el continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
