package main

func insertionSort(src []int) []int {
	n := len(src)
	for i := 1; i < n; i++ {
		tmp := src[i]
		for j := 0; j < i; j++ {
			if src[j] > tmp {
				for k := i; k > j; k-- {
					src[k] = src[k-1]
				}
				src[j] = tmp
				break
			}
		}
	}
	return src
}
