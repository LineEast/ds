package binarytree

func (tree *Tree) InsertAny(a T) {
	node := &Tree{Data: a}
	tree.Insert(node)
}

func (tree *Tree) Insert(node *Tree) {
	if tree.Data == nil {
		tree.Data = node
	}
}
