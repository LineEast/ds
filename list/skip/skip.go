package skip

import (
	"math/rand"
	"time"
)

type (
	SkipList[T any] struct {
		mainList *doublyOrdered[doublyOrdered[T]]

		maxLevel uint
		rand     *rand.Rand

		compare func(a, b T) bool
	}
)

func New[T any](compare func(a, b T) bool, maxlvl uint) *SkipList[T] {
	source := rand.NewSource(time.Now().UnixNano())
	if maxlvl == 0 {
		maxlvl = 32
	}

	l := NewDO[doublyOrdered[T]]()
	for i := 1; i <= int(maxlvl); i++ {
		ll := NewDO[T]()
		ll.level = uint(i)

		l.Push(*ll)
	}

	return &SkipList[T]{
		mainList: l,
		maxLevel: maxlvl,
		rand:     rand.New(source),
		compare:  compare,
	}
}

func (l *SkipList[T]) Push(body T) (node *node[T]) {
	node.body = body

	if l.mainList.head.body.head == nil { // if it is the first node in skip list
		l.mainList.head.body.head = node
		l.mainList.head.body.tail = node
	} else if l.compare(l.mainList.head.body.tail.body, body) {
		node.prev = l.mainList.head.body.tail
		l.mainList.head.body.tail.next = node
		l.mainList.head.body.tail = node
	} else {
	mainList:
		for mainListNode := l.mainList.tail; mainListNode != nil; mainListNode = mainListNode.prev {
			// subList:
			for subListNode := mainListNode.body.head; subListNode != nil; {
				if l.compare(subListNode.body, body) {
					continue mainList
				} else if l.compare(body, subListNode.body) {
					if subListNode.down == nil {
						mainListNode.body.Push(body)
					}

					if subListNode.next == nil || (subListNode.next != nil && l.compare(subListNode.next.body, body)) {
						subListNode = subListNode.down
						mainListNode = mainListNode.down
					} else {
						subListNode = subListNode.next
					}
				}
			}
		}
	}

	return
}

// // if l.compare(subListNode.body, body) { // if slN.body > body
// if subListNode.next == nil {
// 	continue mainList
// } else if  {

// } else if l.compare(body, subListNode.body) && subListNode.down != nil { // if body > sbL.body
// 	continue subList
// }

// if l.compare(subListNode.body, body) { // if slN.body > body
// 	continue mainlist
// } else if l.compare(body, subListNode.body) && subListNode.down != nil { // if body > slN.body
// 	mainListNode = mainListNode.down
// 	subListNode = subListNode.prev.down

// 	continue
// }
