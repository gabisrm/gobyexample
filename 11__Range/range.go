package main

import "fmt"

func main() {

	//recordar que la diferencia entre array y slice es que en la declaración del array hay que indicar el número de elementos,y en slices (como aquí) no
	nums := []int{2, 3, 4}
	sum := 0

	//range itera sobre slices o arrays, o tipos secuenciales
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range devuelve tanto el índice (primer valor del return) como el valor sobre el cual está iterando (segundo return)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//también sobre maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//a través de los keys solamente
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//a través de strings
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
