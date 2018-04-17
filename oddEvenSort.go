package main

func oddEvenSort(src []int) []int {
	for {
		changed := false
		for i := 0; i+1 < len(src); i += 2 {
			if src[i] > src[i+1] {
				changed = true
				src[i], src[i+1] = src[i+1], src[i]
			}
		}
		for i := 1; i+1 < len(src); i += 2 {
			if src[i] > src[i+1] {
				changed = true
				src[i], src[i+1] = src[i+1], src[i]
			}
		}
		if !changed {
			break
		}
	}
	return src
}
