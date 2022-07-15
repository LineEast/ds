package singlelinkedlist

import "log"

func (list *List) DisplayList() {
	if list.len == 0 {
		DisplayError("Empty List")
	} else {
		current := list.head
		for i := 0; uint(i) < list.len; i++ {
			log.Println(current)
			current = current.next
		}
	}
}

func (list *List) DisplayData(i uint) {
	if list.len == 0 || i == 0 || i > list.len {
		DisplayError("Empty List or Index more than Length")
		return
	} else {
		log.Println(list.NodeAt(i))
	}
}

func (list *List) DisplayLen() { log.Println(list.len) }

func (list *List) Len() uint { return list.len }

func DisplayError(s string) { log.Println(s) }