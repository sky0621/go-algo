package main

import "fmt"

func main() {
	a := []*int{&15, &9, &8, &1, &4, &11, &7, &12, &13, &6, &5, &3, &16, &2, &10, &14}
	fmt.Println(a)
	isort(a)
	fmt.Println(a)
}

func isort(a []*int) {
	for pos := 1; pos < 10; pos++ {
		insert(a, pos, a[pos])
	}
}

func insert(a []*int, pos int, value *int) {
	idx := pos - 1
	for idx >= 0 && *a[idx] > *value {
		a[idx+1] = a[idx]
		idx = idx - 1
	}
	a[idx+1] = value
}
