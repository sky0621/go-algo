package main

import "fmt"

func main() {
	taro := &A{}
	jiro := &A{cnt: 100}

	fmt.Println(taro.upup(1))
	fmt.Println(jiro.downdown(10))
}

type A struct {
	cnt int
}

func (a *A) up(n int) {
	a.cnt = a.cnt + n
}

func (a *A) upup(n int) int {
	if n >= 10 {
		return a.cnt
	}
	a.up(n)
	return a.upup(n + 1)
}

func (a *A) down(n int) {
	a.cnt = a.cnt - n
}

func (a *A) downdown(n int) int {
	if n <= 0 {
		return a.cnt
	}
	a.down(n)
	return a.downdown(n - 1)
}
