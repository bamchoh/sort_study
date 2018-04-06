package main

func shakerSort(src []int) []int {
	min := 0
	max := len(src) - 1
	m := 0
	for {
		for j := min; j < max; j++ {
			if src[j] > src[j+1] {
				m = j
				src[j+1], src[j] = src[j], src[j+1]
			}
		}

		if min == m {
			break
		}

		max = m

		for j := max; j > min; j-- {
			if src[j-1] > src[j] {
				m = j
				src[j-1], src[j] = src[j], src[j-1]
			}
		}

		if max == m {
			break
		}

		min = m
	}
	return src
}
