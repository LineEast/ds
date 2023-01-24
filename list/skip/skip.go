package skip

import (
	"math/rand"
	"time"
)

type (
	SkipList[T any] struct {
		mainList *doublyOrdered[doublyOrdered[T]]

		maxLevel uint
		rand     *boolgen

		compare  func(a, b T) bool
		compare2 func(a, b T) int8
	}

	boolgen struct {
		src       rand.Source
		cache     int64
		remaining int
	}
)

const (
	Equel int8 = iota // 0
	More              // 1
	Less              // 2
)

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func New[T any](compare func(a, b T) bool, maxlvl uint) *SkipList[T] {
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
		rand:     &boolgen{src: rand.NewSource(time.Now().UnixNano())},
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
						mainListNode.body.PushNode(node)
						break mainList
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

	for mainListNode := l.mainList.head.next; mainListNode != nil; mainListNode = mainListNode.next {
		if !l.rand.Bool() {
			break
		}

		mainListNode.body.PushNode(node)
	}

	return
}

func (l *SkipList[T]) Find(body T) (node *node[T]) {
mainList:
	for mainList := l.mainList.tail; mainList != nil; {
		for subList := mainList.body.head; subList != nil; {
			switch l.compare2(subList.body, body) {
			case Equel:
				for node = subList; node.down != nil; node = node.down {
				}
				return
			case More:
				continue mainList
			case Less:
				r := l.compare2(subList.next.body, body)
				if r == Equel {
					for node = subList.next; node.down != nil; node = node.down {
					}
					return
				} else if r == More {
					if subList.down == nil {
						node = subList
						return
					}

					subList = subList.down
					mainList = mainList.down

					continue mainList
				}
			}

			subList = subList.next
		}

		mainList = mainList.prev
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
