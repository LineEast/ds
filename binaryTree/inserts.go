package binarytree

import "log"

func (tree *Tree) InsertAny(a int) {
	node := &Tree{Data: a}
	tree.Insert(node)
}

func (tree *Tree) Insert(node *Tree) {
	if tree == nil {
		log.Println("Empty Tree")
		return
	}
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
		return
	}
	if tree.Data > node.Data {
		if tree.NodeL == nil {
			tree.NodeL = node
			return
		}
		tree.NodeL.Insert(node)
	} else {
		if tree.NodeR == nil {
			tree.NodeR = node
			return
		}
		tree.NodeR.Insert(node)
	}
}
