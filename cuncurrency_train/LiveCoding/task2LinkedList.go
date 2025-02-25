package main

type LinkNode struct {
	next  *LinkNode
	value int
}

func (ln *LinkNode) Print() {
	// TODO: implement method
}

func main() {
	// инициализируем элементы списка
	var n1, n2, n3 = LinkNode{value: 1}, LinkNode{value: 2}, LinkNode{value: 3}
	n1.next, n2.next = &n2, &n3 // и связываем их

	n1.Print()

	// и теперь обратим список в зад
	var prev, next *LinkNode

	//	крутим цикл до тех пор, пока current не станет nil
	for current := &n1; current != nil; {
		next, current.next = current.next, prev
		prev, current = current, next
	}

	n3.Print()
}
// package main

// type LinkNode struct {
// 	next  *LinkNode
// 	value int
// }

// func (l *LinkNode) Print() {
// 	// ставим current указателем на первый элемент, на каждой итерации заменяя его на next
// 	// до тех пор, пока current не станет nil
// 	for current := l; current != nil; current = current.next {
// 		print(current.value)

// 		if current.next != nil {
// 			println(" ->", current.next.value)
// 		}
// 	}

// 	println()
// }

// func main() {
// 	// инициализируем элементы списка
// 	var n1, n2, n3 = LinkNode{value: 1}, LinkNode{value: 2}, LinkNode{value: 3}
// 	n1.next, n2.next = &n2, &n3 // и связываем их

// 	n1.Print()
// 	// 1 -> 2
// 	// 2 -> 3
// 	// 3

// 	// и теперь обратим список в зад
// 	var prev, next *LinkNode

// 	// крутим цикл до тех пор, пока current не станет nil
// 	for current := &n1; current != nil; {
// 		next, current.next = current.next, prev
// 		prev, current = current, next
// 	}

// 	n3.Print()
// 	// 3 -> 2
// 	// 2 -> 1
// 	// 1
// }
