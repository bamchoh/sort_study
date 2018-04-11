package main

func mergeSort(src []int) []int {
	n := len(src)
	if n == 0 || n == 1 {
		return src
	}
	ary1 := src[0 : n/2]
	mary1 := mergeSort(ary1)
	ary2 := src[n/2:]
	mary2 := mergeSort(ary2)
	var ret []int
	i := 0
	j := 0
	for {
		if len(mary1[i:]) == 0 {
			if len(mary2[j:]) == 0 {
				break
			} else {
				ret = append(ret, mary2[j])
				j++
				continue
			}
		}

		if len(mary2[j:]) == 0 {
			ret = append(ret, mary1[i])
			i++
			continue
		}

		if mary1[i] > mary2[j] {
			ret = append(ret, mary2[j])
			j++
		} else {
			ret = append(ret, mary1[i])
			i++
		}
	}
	return ret
}
