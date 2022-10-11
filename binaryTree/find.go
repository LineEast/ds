package binarytree

func (tree *Tree) Find(x int) bool {
	if x == tree.Data {
		return true
	} else {
		if x > tree.Data && tree.NodeR != nil {
			if f := tree.NodeR.Find(x); f {
				return true
			}
		} else if tree.NodeL != nil {
			if f := tree.NodeL.Find(x); f {
				return true
			}
		}
	}
	return false
}
