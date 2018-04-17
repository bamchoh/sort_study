package main

import (
	"runtime"
	"sync"
)

func oddEvenSort2(src []int) []int {
	for {
		changed := false
		wg := &sync.WaitGroup{}
		cpu := runtime.NumCPU()
		for x := 0; x < cpu; x++ {
			wg.Add(1)
			go func(j int) {
				for i := j * 2; i+1 < len(src); i += 2 * cpu {
					if src[i] > src[i+1] {
						changed = true
						src[i], src[i+1] = src[i+1], src[i]
					}
				}
				wg.Done()
			}(x)
		}
		wg.Wait()
		for x := 0; x < cpu; x++ {
			wg.Add(1)
			go func(j int) {
				for i := j*2 + 1; i+1 < len(src); i += 2 * cpu {
					if src[i] > src[i+1] {
						changed = true
						src[i], src[i+1] = src[i+1], src[i]
					}
				}
				wg.Done()
			}(x)
		}
		wg.Wait()
		if !changed {
			break
		}
	}
	return src
}
