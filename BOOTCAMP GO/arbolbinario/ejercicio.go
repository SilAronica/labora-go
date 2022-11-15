package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func Walk(t *tree.Tree, ch chan int){
	tree.New(k)
	go Walk(tree.New(1), ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	Same(tree.New(1), tree.New(1))
	Same(tree.New(1), 
}

func main() {
}
