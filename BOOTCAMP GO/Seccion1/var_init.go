package main

import (
    "fmt"
)

func main() {

    var i, j, k = 1, 2, 3

    var (
        name       = "John Doe"
        occupation = "gardener"
    )

    fmt.Println(i, j, k)
    fmt.Printf("%s is a %s\\n", name, occupation)
}