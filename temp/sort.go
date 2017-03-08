package main

import "fmt"

func main() {
	pgs := []string{"Go", "Java", "Ruby", "Python", "C", "PHP", "C++", "JavaScript", "Elixir", "C#", "Scala"}
	proc(pgs, 1)
}

func proc(ar []string, i int) {
	if i >= len(ar)-1 {
		fmt.Println("End")
		return
	}
	org := ar[i]
	bef := ar[i-1]

	if len(bef) < len(org) {
		fmt.Println(org)
	}
	proc(ar, i+1)
}
