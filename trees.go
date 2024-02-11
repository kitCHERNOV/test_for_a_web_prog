package main

import (
	//"fmt"
	"errors"
	"fmt"
	"log"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == value {
		return errors.New("This node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}


func main() {
	Tree := TreeNode{}
	var err error
	err = Tree.Insert(13)
	if err!=nil {
		log.Fatal(err)
	}
	err = Tree.Insert(1)
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println(Tree.FindMax())
	
	
	
}