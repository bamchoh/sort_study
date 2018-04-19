package main

import "math"

func ssSort2(src []int, l, r int, fn func(int, int) bool) {
	if l >= r {
		return
	}
	i, j := l, r-1
	for i <= j {
		for i <= r && fn(src[r], src[i]) {
			i++
		}
		for l <= j && fn(src[j], src[r]) {
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
	ssSort2(src, l, i-1, fn)
	ssSort2(src, i+1, r, fn)
}

func ssSort1(src []int, fn func(int, int) bool) {
	ssSort2(src, 0, len(src)-1, fn)
}

func shearSort(src []int) []int {
	n := len(src)
	m := int(math.Sqrt(float64(n)))
	cnt := 2 * (int(math.Log2(float64(n))) + 1)
	for i := 0; i < int(cnt); i++ {
		for y := 0; y <= n/m; y++ {
			var fn func(int, int) bool
			if y%2 == 0 {
				fn = func(a, b int) bool {
					return a > b
				}
			} else {
				fn = func(a, b int) bool {
					return a < b
				}
			}
			var l, r int
			if (y+1)*m > n {
				l = y * m
				r = n
			} else {
				l = y * m
				r = (y + 1) * m
			}
			ssSort1(src[l:r], fn)
		}

		for x := 0; x < m; x++ {
			for y := 0; (y+1)*m+x < n; y++ {
				for z := 0; (z+1)*m+x < n; z++ {
					j := z*m + x
					if src[j] > src[j+m] {
						src[j], src[j+m] = src[j+m], src[j]
					}
				}
			}
		}
	}

	for y := 0; y <= n/m; y++ {
		var l, r int
		if (y+1)*m > n {
			l = y * m
			r = n
		} else {
			l = y * m
			r = (y + 1) * m
		}
		ssSort1(src[l:r], func(a, b int) bool {
			return a > b
		})
	}

	return src
}
