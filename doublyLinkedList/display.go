package doublylinkedlist

import "log"
import "fmt"

func (list *List[T]) DisplayList() {
	if list.len == 0 {
		DisplayError("Empty List")
	} else {
		current := list.head
		for i := 0; uint(i) < list.len; i++ {
			fmt.Println("Data:", current.Data, "\tPrev:", current.prev, "\tNext:", current.next)
			current = current.next
		}
		fmt.Println("-------------------")
	}
}

func (list *List[T]) DisplayData(i uint) {
	if list.len == 0 || i == 0 || i > list.len {
		DisplayError("Empty List or Index more than Length")
		return
	} else {
		log.Println(list.NodeAt(i))
	}
}

func (list *List[T]) DisplayLen() { log.Println(list.len) }

func (list *List[T]) Len() uint { return list.len }

func DisplayError(s string) { log.Println(s) }
