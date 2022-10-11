package data

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	binarytree "github.com/go-asphyxia/data/binaryTree"
	doublylinkedlist "github.com/go-asphyxia/data/doublyLinkedList"
	singlelinkedlist "github.com/go-asphyxia/data/singleLinkedList"
)

func singleLinkedListExample() {
	list := singlelinkedlist.Init[int]()
	array := []int{1, 2, 3, 4, 5}

	for _, el := range array {
		list.InsertFirstAny(el)
	}

	list.DisplayList()
}

func doublyLinkedListExample() {
	dlist := doublylinkedlist.Init[int]()
	array := []int{1, 2, 3, 4, 5}

	for _, el := range array {
		dlist.InsertLastAny(el)
	}

	dlist.DisplayList()
	dlist.InsertIndexAny(65464, 2)
	dlist.DisplayList()
}

func binarytreeExample() {
	tree := binarytree.Init()

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for _, el := range array {
		tree.InsertAny(el)
	}

	spew.Dump(tree)
	tree.Sort()
	spew.Dump(tree)
	fmt.Println(tree.Find(15))
}

func main() {
	// singleLinkedListExample()
	// doublyLinkedListExample()
	binarytreeExample()
}