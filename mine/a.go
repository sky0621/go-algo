package main

import (
	"fmt"
	"time"
)

func main() {
	cache := []map[int]interface{}
	st := time.Now()
	for i := 1; i < 1000; i++ {
		for j := 1000; j > 1; j-- {
			if
			fmt.Printf("%d, ", i*j)
		}
	}
	fmt.Println()
	fmt.Println(time.Since(st))
}
