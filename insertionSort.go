package main

func insertionSort(src []int) []int {
	n := len(src)
	for i := 1; i < n; i++ {
		e := src[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if src[j] <= e {
				break
			}
		}

		for k := i; k > j+1; k-- {
			src[k] = src[k-1]
		}
		src[j+1] = e
	}
	return src
}
