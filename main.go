package main

import (
	"fmt"
)

type Record struct {
	key   int
	value int
}

type Tree struct {
	value Record
	left  *Tree
	right *Tree
}

func printTree(t *Tree) {
	if t.left != nil {
		printTree(t.left)
	}
	if &t.value != nil {
		fmt.Println(t.value.value)
	}
	if t.right != nil {
		printTree(t.right)
	}
}

func appendTree(t *Tree, r Record) {
	newTree := Tree{value: r, left: nil, right: nil}
	if r.value > t.value.value {
		if t.right == nil {
			t.right = &newTree
		} else {
			appendTree(t.right, r)
		}
	} else {
		if t.left == nil {
			t.left = &newTree
		} else {
			appendTree(t.left, r)
		}
	}
}

func main() {
	r1 := Record{key: 1, value: 2}
	r2 := Record{key: 2, value: 4}
	r3 := Record{key: 3, value: 1}
	r4 := Record{key: 4, value: 6}
	r5 := Record{key: 5, value: 0}

	var t1 Tree = Tree{value: r1, left: nil, right: nil}
	appendTree(&t1, r2)
	appendTree(&t1, r3)
	appendTree(&t1, r4)
	appendTree(&t1, r5)
	printTree(&t1)
}
