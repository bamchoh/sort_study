package main

func combSort(src []int) []int {
	n := len(src)
	h := n
	for {
		if h != 1 {
			h = int(float64(h) / 1.3)
		}

		change := false
		for i := 0; i+h < n; i++ {
			if src[i] > src[i+h] {
				change = true
				src[i], src[i+h] = src[i+h], src[i]
			}
		}

		if !change {
			break
		}
	}
	return src
}
