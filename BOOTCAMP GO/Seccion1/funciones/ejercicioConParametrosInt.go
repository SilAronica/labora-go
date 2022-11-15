package main

import "fmt"

func main() {
	addNumbers(1, 2, 3)
}

func addNumbers(x, y, z int) {
	a := x + y
	b := x + z
	c := y + z
	fmt.Println(a, b, c)
}