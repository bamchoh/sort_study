package main

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var size = flag.Int("size", 20, "size of array")
var repeat = flag.Int("repeat", 1, "number of repeat testing")

func testMain(t *testing.T, sort func([]int) []int) {
	rand.Seed(time.Now().UnixNano())
	var want []int
	for i := 0; i < *size; i++ {
		want = append(want, i)
	}

	for x := 0; x < *repeat; x++ {
		t.Run(fmt.Sprintf("size:%v,x:%v", *size, x), func(t *testing.T) {
			var tt = rand.Perm(*size)
			got := sort(tt)

			if len(got) != len(want) {
				t.Fatalf("length is different(got:%v,want:%v)", got, want)
			}

			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Fatalf("index [%v] is different:\ngot :%v\nwant:%v", i, got, want)
				}
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testMain(t, bubbleSort)
}

func TestShakerSort(t *testing.T) {
	testMain(t, shakerSort)
}

func TestCombSort(t *testing.T) {
	testMain(t, combSort)
}

func TestGnomeSort(t *testing.T) {
	testMain(t, gnomeSort)
}
