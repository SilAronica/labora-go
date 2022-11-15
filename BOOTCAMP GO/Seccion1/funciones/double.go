package main

import "fmt"

func main() {
	result := double(3)
	fmt.Println(result)
}

func double(x int) int {
	y := x * 2
	return y
}