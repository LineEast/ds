package binarytree

import (
	"fmt"
	"sort"
)

func (tree *Tree) Sort() {
	array := tree.writeToArray()
	sort.Ints(array)
	tree.NodeL = nil
	tree.NodeR = nil
	tree.Data = 0
	i := len(array)
	for len(array) > 1 {
		i = i / 2
		tree.InsertAny(array[i])
		array = append(array[:i], array[i+1:]...) // delete
		if i == 1 {
			i = len(array)
		}

	}
	tree.InsertAny(array[0])
	fmt.Println("Sorting Complete")
}

func (tree *Tree) writeToArray() (array []int) {
	if tree == nil {
		return
	}
	array = append(array, tree.NodeL.writeToArray()...)
	array = append(array, tree.Data)
	array = append(array, tree.NodeR.writeToArray()...)
	return array
}
