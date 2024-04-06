package main

import "fmt"

func add(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3
	fmt.Println("x = ", x)

	b := add(&x)
	fmt.Println("x+1 = ", b) // debe imprimir "x+1 = 4"
	fmt.Println("x = ", x)   // debe imprimir "x = 4"

}
