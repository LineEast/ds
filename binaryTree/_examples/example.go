package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-asphyxia/data/binaryTree"
)

func main() {
	tree := binarytree.Init()

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < len(array); i++ {
		tree.InsertAny(array[i])
	}
	spew.Dump(tree)
	tree.Sort()
	spew.Dump(tree)
	fmt.Println(tree.Find(15))
}
