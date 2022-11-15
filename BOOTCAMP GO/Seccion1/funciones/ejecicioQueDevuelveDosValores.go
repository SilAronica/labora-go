package main

import "fmt"

func addAndMultiply(i, j int)(int, int) {
	return i + j, i * j
}

func main() {
	addResult, mulResult := addAndMultiply(2,3)
	fmt.Println("suma de 2+3 es ", addResult)
	fmt.Println("multiplicación de 2*3 es ", mulResult)
}