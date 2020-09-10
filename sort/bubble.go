package sort

func Bubble(a []int, args ...bool) {

	des := false
	if len(args) > 0 {
		des = args[0]
	}
	bubble(a, des)
}

func bubble(a []int, des bool) {

	var n int
	for j := len(a); j > 0; j-- {
		for i := 0; i+1 < j; i++ {
			swap := false
			if des && a[i+1] > a[i] {
				swap = true
			} else if !des && a[i+1] < a[i] {
				swap = true
			}
			if swap {
				n = a[i]
				a[i] = a[i+1]
				a[i+1] = n
			}
		}
	}
}
