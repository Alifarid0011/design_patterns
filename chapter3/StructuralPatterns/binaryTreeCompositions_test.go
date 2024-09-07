package main

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := Tree{
		LeafValue: 10,
		Right:     &Tree{15, &Tree{17, nil, nil}, nil},
		Left:      &Tree{0, nil, nil},
	}
	fmt.Println(root.Right.Right.LeafValue)
}
