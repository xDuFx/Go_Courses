package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func New(q int) *LinkedList {
	First := Node{}
	NewM := &First
	for i := 1; i < q; i++ {
		NewM.Next = &Node{}
		NewM = NewM.Next
	}
	Link := LinkedList{}
	Link.Head = &First
	return &Link
}

func (l *LinkedList) Add(val int) {
	NewM := l.Head
	for i := 0; i >= 0; {
		if NewM.Next == nil {
			NewM.Next = &Node{val, nil}
			break
		}
		NewM = NewM.Next
	}
}
func (l *LinkedList) Pop() {
	NewM := l.Head
	Previ := l.Head
	for i := 0; i < 1; {
		if NewM.Next == nil {
			Previ.Next = &Node{0, nil}
			break
		}
		Previ = NewM
		NewM = NewM.Next
	}
}

func (l *LinkedList) Size() int {
	NewM := l.Head
	i := 0
	for ; i >= 0; i++ {
		if NewM.Next == nil {
			break
		}
		NewM = NewM.Next

	}
	return i + 1
}

func (l *LinkedList) At(pos int) int {
	NewM := l.Head
	at := 0
	for i := 0; i < pos; i++ {
		if NewM.Next == nil && i != pos-1 {
			fmt.Println("Выход за границы")
			return NewM.Value
		}
		if i == pos-1 {
			at = NewM.Value
			break
		}
		NewM = NewM.Next
	}
	return at
}

func (l *LinkedList) DeleteFrom(pos int) {
	NewM := l.Head
	Previ := l.Head
	for i := 0; i < pos; i++ {
		if NewM.Next == nil && i != pos-1 {
			fmt.Println("Выход за границы")
			Previ.Next = &Node{0, nil}
			break
		} else if i == pos-1 {
			Previ.Next = &Node{0, nil}
			break
		}
		Previ = NewM
		NewM = NewM.Next
	}
}

func (l *LinkedList) UpdateAt(val, pos int) {
	NewM := l.Head
	for i := 0; i < pos; i++ {
		if NewM.Next == nil && i != pos-1 {
			fmt.Println("Выход за границы")
			NewM.Value = val
			break
		} else if i == pos-1 {
			NewM.Value = val
			break
		}
		NewM = NewM.Next
	}
}

func NewFromSlice(s []int) *LinkedList {
	First := Node{s[0], nil}
	NewM := &First
	for i := 1; i < len(s); i++ {
		NewM.Next = &Node{s[i], nil}
		NewM = NewM.Next
	}
	Link := LinkedList{}
	Link.Head = &First
	return &Link
}

func main() {
	s := []int{3, 4, 5}
	l := NewFromSlice(s)
	p := New(3)
	fmt.Println(p.Size())
	p.UpdateAt(3,3)
	fmt.Println(p.At(3))
	l.Add(4)
	fmt.Println(l.Size())
	l.Pop()
	fmt.Println(l.Size())
	l.DeleteFrom(1)
}
