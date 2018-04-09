package main

import "fmt"

type binaryTree struct {
	Value int
	Left  *binaryTree
	Right *binaryTree
}

func (tree *binaryTree) Print(dup int) {
	fmt.Println("val:", tree.Value)
	if tree.Left != nil {
		for i := 0; i < dup; i++ {
			fmt.Print(" ")
		}
		fmt.Print("Left:")
		tree.Left.Print(dup + 1)
	} else {
		for i := 0; i < dup; i++ {
			fmt.Print(" ")
		}
		fmt.Print("Left:")
		fmt.Print("<nil>\n")
	}

	if tree.Right != nil {
		for i := 0; i < dup; i++ {
			fmt.Print(" ")
		}
		fmt.Print("Right:")
		tree.Right.Print(dup + 1)
	} else {
		for i := 0; i < dup; i++ {
			fmt.Print(" ")
		}
		fmt.Print("Right:")
		fmt.Print("<nil>\n")
	}
}

func (tree *binaryTree) Insert(val int) {
	if tree.Value > val {
		if tree.Left == nil {
			tree.Left = newNode(val)
		} else {
			tree.Left.Insert(val)
		}
	} else {
		if tree.Right == nil {
			tree.Right = newNode(val)
		} else {
			tree.Right.Insert(val)
		}
	}
}

func (tree *binaryTree) ToList() []int {
	var src []int

	if tree.Left != nil {
		src = tree.Left.ToList()
	}

	src = append(src, tree.Value)

	if tree.Right != nil {
		src2 := tree.Right.ToList()
		src = append(src, src2...)
	}

	return src
}

func newNode(val int) *binaryTree {
	return &binaryTree{
		Value: val,
	}
}

func binarySearchSort(src []int) []int {
	var root *binaryTree
	for _, val := range src {
		if root == nil {
			root = newNode(val)
		} else {
			root.Insert(val)
		}
	}
	return root.ToList()
}
