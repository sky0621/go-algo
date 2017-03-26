package main

import (
	"fmt"

	"github.com/sky0621/go-algo/util"
)

func main() {
	targets := util.RandomIntsNoDup(100)
	fmt.Println(targets)

	for ix := 0; ix < len(targets)-1; ix++ {
		for idx := 0; idx < len(targets)-1; idx++ {
			if targets[idx] > targets[idx+1] {
				tmp := targets[idx+1]
				targets[idx+1] = targets[idx]
				targets[idx] = tmp
			}
		}
	}

	fmt.Println(targets)
}
