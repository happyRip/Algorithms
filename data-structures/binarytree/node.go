package binarytree

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

func NewNode(v int) *node {
	return &node{v, nil, nil, nil}
}
