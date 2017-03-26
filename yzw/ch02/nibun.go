package main

import (
	"fmt"

	"github.com/sky0621/go-algo/util"
)

func main() {
	targets := util.RandomIntsNoDup(10)
	util.Bsort(targets)
	fmt.Println(targets)

	x := 7

	fmt.Println(binSearch(targets, x, targets[len(targets)/2-1]))
}

func binSearch(targets []int, x, pos int) int {
	if len(targets) < 1 {
		return -1
	}
	if targets[pos] == x {
		return pos
	}
	if targets[pos] > x {
		return binSearch(targets[0:pos], x, pos/2-1)
	} else {
		return binSearch(targets[pos:len(targets)], x, len(targets)/2-1)
	}
}
