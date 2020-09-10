package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IsSorted(a []int) bool {

	if len(a) <= 1 {
		return true
	}

	n := a[0]
	i := 1
	for ; i < len(a) && a[i] == n; i++ {
	}
	if i == len(a)-1 {
		return true
	}
	des := a[i] < n
	for ; i < len(a); i++ {
		if des && a[i] > n {
			return false
		} else if !des && a[i] < n {
			return false
		}
		n = a[i]
	}
	return true
}

func RandList(size int, max int) []int {

	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(max)
	}
	return a
}
