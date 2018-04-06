package main

func bubbleSort(src []int) []int {
	for i := 0; i < len(src)-1; i++ {
		for j := 0; j < len(src)-1; j++ {
			if src[j] > src[j+1] {
				src[j+1], src[j] = src[j], src[j+1]
			}
		}
	}
	return src
}
