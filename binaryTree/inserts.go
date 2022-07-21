package binarytree

import "log"

func (tree *Tree) InsertAny(a int) {
	node := &Tree{Data: a}
	tree.Insert(node)
}

func (tree *Tree) Insert(node *Tree) {
	if node.Data == 0 {
		log.Println("Insert can't be zero.")
		return
	}
	if tree.Data == node.Data {
		log.Println("Value in Tree already exist. Skipping Insert. ¯\\_(ツ)_/¯")
		return
	}
	if tree.Data == 0 {
		tree.Data = node.Data
	} else if tree.Data < node.Data {
		tree.NodeL.Insert(node)
	} else {
		tree.NodeR.Insert(node)
	}
}