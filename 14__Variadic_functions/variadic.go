package main

import "fmt"

//como en javascript el rest (es6), pero aquí se pone al tipo
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	//estas funciones se pueden llamar con argumentos de entrada independientes
	sum(1, 2)
	sum(1, 2, 3, 4, 5)

	//o también pasandole un slice
	nums := []int{1, 2, 3, 4, 5}
	//si el slice ya tiene multiples argumentos, se le pasa con este formato
	sum(nums...)

}
