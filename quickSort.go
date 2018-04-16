package main

func sort(src []int, l, r int) {
	if l >= r {
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
	src[r], src[i] = src[i], src[r]
	sort(src, l, i-1)
	sort(src, i+1, r)
}

func quickSort(src []int) []int {
	sort(src, 0, len(src)-1)
	return src
}
