package main

import (
	"fmt"
	"os"
)

func main() {
	saiki(105)
}

func saiki(n int) {
	if n <= 0 {
		fmt.Println("ENd")
		os.Exit(0)
	}
	fmt.Println(n * n)
	saiki(n - 1)
}
