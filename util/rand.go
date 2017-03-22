package util

import (
	"math/rand"
	"time"
)

func RandomInts(len int) []int {
	ints := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		ints = append(ints, rand.Intn(len))
	}
	return ints
}
