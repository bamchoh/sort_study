package main

func shellSort(src []int) []int {
	n := len(src)
	for h := n / 2; h > 0; h /= 2 {
		for x := 0; x < h; x++ {
			for i := h + x; i < n; i += h {
				e := src[i]
				var j int
				for j = i - h; j >= x; j -= h {
					if src[j] <= e {
						break
					}
				}

				for k := i; k > j+h; k -= h {
					src[k] = src[k-h]
				}
				src[j+h] = e
			}
		}
	}
	return src
}
