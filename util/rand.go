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

func RandomIntsNoDup(l int) []int {
	ints := []int{}
	rand.Seed(time.Now().UnixNano())
	for len(ints) < l {
		ri := rand.Intn(l)
		if IsContain(ri, ints) {
			continue
		}
		ints = append(ints, ri)
	}
	return ints
}

func IsContain(t int, nos []int) bool {
	for _, no := range nos {
		if no == t {
			return true
		}
	}
	return false
}
