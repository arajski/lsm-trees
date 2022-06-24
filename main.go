package main

import (
    "math/rand"
    "lsm-trees/lsmtree"
)

func main() {
    var lsmTree *lsmtree.Tree = lsmtree.CreateLSM()
    for i:=0; i < 20; i++ {
        r := lsmtree.Record{Key: rand.Int() % 10, Value: rand.Int()}
        lsmtree.Push(lsmTree,&r)
    }
    lsmtree.Print(lsmTree)
}
