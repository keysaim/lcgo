package sort

func Quick(a []int, args ...bool) {

	des := false
	if len(args) > 0 {
		des = args[0]
	}
	quick(a, des)
}

func quick(a []int, des bool) {

	if len(a) <= 1 {
		return
	}

	i := 0
	j := len(a) - 1
	var n int
	for {
		if des {
			for ; j > i && a[j] < a[0]; j-- {
			}
			for ; i < j && a[i] >= a[0]; i++ {
			}
		} else {
			for ; j > i && a[j] > a[0]; j-- {
			}
			for ; i < j && a[i] <= a[0]; i++ {
			}
		}
		if i == j {
			if j > 0 {
				n = a[0]
				a[0] = a[j]
				a[j] = n
			}
			quick(a[:j], des)
			quick(a[j+1:], des)
			return
		} else {
			n = a[i]
			a[i] = a[j]
			a[j] = n
		}
	}
}
