package main

func selectionSort(src []int) []int {
	n := len(src)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
	return src
}
