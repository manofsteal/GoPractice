package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }

func Walk(tr *tree.Tree, vals *[]int) {

	if tr == nil {
		return
	}

	Walk(tr.Left)
	// fmt.Println("val ", tr.Value)
	vals = append(vals, tr.Value)
	Walk(tr.Right)
}

func Same(tr1 *tree.Tree, tr2 *tree.Tree)

func main() {
	fmt.Println("Program begin")

	Walk(tree.New(1))

}
