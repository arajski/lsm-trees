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

func Print(t *Tree) {
    if t.value == nil {
        return
    }
    if t.left != nil {
        Print(t.left)
    }
    if t.right != nil {
        Print(t.right)
    }
    fmt.Println(t.value.Key,":", t.value.Value)
}
func Size(t *Tree) int{
  var size int = 0
  if t.left != nil {
    size += Size(t.left)
  }
  if t.right != nil {
      size += Size(t.right)
  }
  if t.value != nil {
      size += 1
  }
  return size 
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
