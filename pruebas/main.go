package main

import (
	"fmt"
)

func math() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	x := math()
	fmt.Println(x())
	fmt.Println(x())

}
