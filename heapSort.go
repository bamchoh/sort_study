package main

func upheap(a []int, n int) []int {
	if n < 1 {
		return a
	}
	pn := (n+1)/2 - 1
	if a[pn] > a[n] {
		a[pn], a[n] = a[n], a[pn]
		a = upheap(a, pn)
	}
	return a
}

func insert(a []int, v int) []int {
	a = append(a, v)
	return upheap(a, len(a)-1)
}

func downheap(a []int, k int) []int {
	v := a[k]
	var j int
	for ; k < len(a)/2; k = j {
		j = 2*(k+1) - 1
		if j < len(a)-1 && a[j] > a[j+1] {
			j++
		}
		if v <= a[j] {
			break
		}
		a[k] = a[j]
	}
	a[k] = v
	return a
}

func remove(a []int) (int, []int) {
	v := a[0]
	a[0] = a[len(a)-1]
	if len(a) == 1 {
		return v, make([]int, 0)
	}
	a = downheap(a[0:len(a)-1], 0)
	return v, a
}

func heapSort(src []int) []int {
	var a []int
	for _, v := range src {
		a = insert(a, v)
	}
	var v int
	for i := 0; len(a) != 0; i++ {
		v, a = remove(a)
		src[i] = v
	}
	return src
}
