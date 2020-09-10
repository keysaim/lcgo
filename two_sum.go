package main

import (
	"fmt"

	"./sort"
)

func twosum1(a []int, sum int) (int, int) {

	m := make(map[int]int, len(a))
	for i, v := range a {
		if j, ok := m[sum-v]; ok {
			return i, j
		}
		m[v] = i
	}
	return -1, -1
}

func twosum2(a []int, sum int) (int, int) {

	sort.Quick(a)
	i := 0
	j := len(a) - 1
	for i < j {
		s := a[i] + a[j]
		if s == sum {
			return i, j
		}
		if s > sum {
			j--
		} else {
			i++
		}
	}
	return -1, -1
}

func main() {

	sum := 100
	a := sort.RandList(100, sum)
	i, j := twosum1(a, sum)
	if i < 0 {
		fmt.Printf("Not found sum %d, a: %v\n", sum, a)
	} else {
		fmt.Printf("%d = %d(i%d) + %d(i%d), a: %v\n", sum, a[i], i, a[j], j, a)
	}

	i, j = twosum2(a, sum)
	if i < 0 {
		fmt.Printf("Not found sum %d, a: %v\n", sum, a)
	} else {
		fmt.Printf("%d = %d(i%d) + %d(i%d), a: %v\n", sum, a[i], i, a[j], j, a)
	}

}
