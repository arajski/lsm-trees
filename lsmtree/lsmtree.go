package lsmtree

import(
    "fmt"
)

type Record struct {
    Key   int
    Value int
}

type Tree struct {
    value *Record
    left  *Tree
    right *Tree
}

func Print(t *Tree) *Tree {
    if t.left != nil {
        Print(t.left)
    }
    if &t.value != nil {
        fmt.Println(t.value.Key,":", t.value.Value)
    }
    if t.right != nil {
        Print(t.right)
    }

    return t
}

func Push(t *Tree, r *Record) *Record {
    if t.value == nil {
           t.value = r 
           return r
    }
    if r.Key > t.value.Key {
        if t.right == nil {
            newTree := Tree{value: r, left: nil, right: nil}
            t.right = &newTree
        } else {
            Push(t.right, r)
        }
    } else {
        if t.left == nil {
            newTree := Tree{value: r, left: nil, right: nil}
            t.left = &newTree
        } else {
            Push(t.left, r)
        }
    }
    return r
}

func CreateLSM() *Tree {
    var tree Tree = Tree{value: nil, left: nil, right: nil}
    return &tree 
}
