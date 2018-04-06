package main

func gnomeSort(src []int) []int {
	n := len(src) - 1
	for i := 0; i < n; i++ {
		if src[i] > src[i+1] {
			src[i], src[i+1] = src[i+1], src[i]
			for j := i; j > 0; j-- {
				if src[j-1] > src[j] {
					src[j-1], src[j] = src[j], src[j-1]
				} else {
					break
				}
			}
		}
	}
	return src
}
