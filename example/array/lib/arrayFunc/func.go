package array

import (
	"math/rand/v2"
)

func RandomFill(src[]int) []int {
	for i := range src {
		src[i] = rand.IntN(100)
	}
	return (src)
}

func RangeFill(src[]int, start int,) []int {
	for i := 0 ; i < len(src) ; i++{
		src[i] = start + i
	}
	return (src)
}
