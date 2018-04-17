package main

import "math"

func sort1(src []int, l, r, d int) {
	if l >= r {
		return
	}
	if d == 0 {
		heapSort(src[l : r+1])
		return
	}
	i, j := l, r-1
	for i <= j {
		for i <= r && src[i] < src[r] {
			i++
		}
		for l <= j && src[j] >= src[r] {
			j--
		}
		if i > j {
			break
		}
		src[i], src[j] = src[j], src[i]
		i++
		j--
	}
	src[i], src[r] = src[r], src[i]
	sort1(src, l, i-1, d-1)
	sort1(src, i+1, r, d-1)
}

func introSort(src []int) []int {
	d := math.Log2(float64(len(src)))
	sort1(src, 0, len(src)-1, int(2*d))
	return src
}
