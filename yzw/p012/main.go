package main

import "fmt"

func main() {
	ml := &MyList{
		Val: 1,
		Child: &MyList{
			Val: 2,
			Child: &MyList{
				Val: 3,
				Child: &MyList{
					Val:   4,
					Child: nil,
				},
			},
		},
	}
	ml.doOnce()
}

type MyList struct {
	Val   int
	Child *MyList
}

func (m *MyList) hasChild() bool {
	return m.Child != nil
}

func (m *MyList) doOnce() {
	fmt.Printf("Val:%d\n", m.Val)
	if m.hasChild() {
		m.Child.doOnce()
	}
}
