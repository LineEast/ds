package binarytree

import "fmt"

func (tree *Tree) DisplayTree() {
	if tree == nil {
		return
	}
	tree.NodeL.DisplayTree()
	fmt.Println(tree.Data)
	tree.NodeR.DisplayTree()
}
