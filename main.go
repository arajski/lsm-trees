package main

import (
    "fmt"
    "time"
    "os"
    "math/rand"
    "lsm-trees/lsmtree"
)

func createLog(tree *lsmtree.Tree){
    r := lsmtree.Record{Key: rand.Int() % 10, Value: rand.Int()}
    fmt.Println("Inserting record ", &r)
    lsmtree.Push(tree, &r)
    fmt.Println("Tree size : ", lsmtree.Size(tree))
}

func dump(tree *lsmtree.Tree) {
  var filename string = fmt.Sprintf("%d.txt", time.Now().Unix())
  fmt.Println("Dumping tree to ", filename)
  lsmtree.Print(tree)
  os.Create(filename) 
}

func main() {
    var tree *lsmtree.Tree = lsmtree.CreateLSM()
    for {
       if lsmtree.Size(tree) == 5 {
           go dump(tree)
           tree = lsmtree.CreateLSM()
       }
       go createLog(tree)
       time.Sleep(time.Second * 1)
    }
}
