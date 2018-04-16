package main

func upheap(a []int, n int) {
	if n < 1 {
		return
	}
	pn := (n+1)/2 - 1
	if a[pn] < a[n] {
		a[pn], a[n] = a[n], a[pn]
		upheap(a, pn)
	}
}

func insert(a []int, n int) {
	upheap(a, n)
}

func downheap(a []int, k int, n int) {
	v := a[k]
	var j int
	for ; k < (n+1)/2; k = j {
		j = 2*(k+1) - 1
		if j <= n-1 && a[j] < a[j+1] {
			j++
		}
		if v >= a[j] {
			break
		}
		a[k] = a[j]
	}
	a[k] = v
}

func remove(a []int, n int) {
	if n == 0 {
		return
	}

	a[0], a[n] = a[n], a[0]
	downheap(a, 0, n-1)
}

func heapSort(src []int) []int {
	for i := 0; i < len(src); i++ {
		insert(src, i)
	}
	for i := len(src) - 1; i >= 0; i-- {
		remove(src, i)
	}
	return src
}
