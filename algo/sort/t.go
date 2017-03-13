package main

import "fmt"

func main() {
	a := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	fmt.Println(a)
	fmt.Println("=========================================--")
	for m := 0; m < len(a)-1; m++ {
		for n := 0; n < len(a)-1; n++ {
			if a[n] > a[n+1] {
				tmp := a[n]
				a[n] = a[n+1]
				a[n+1] = tmp
				fmt.Println(a)
			}
		}
	}
	fmt.Println("=========================================--")
	fmt.Println(a)
}
