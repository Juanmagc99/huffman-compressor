package tree

type BaseNode interface {
	IsLeaf() bool
	Weight() int
}

type LeafNode struct {
	Character rune
	RepNumber int
}

type InternalNode struct {
	LeftNode  BaseNode
	RightNode BaseNode
	RepNumber int
}

func (n LeafNode) IsLeaf() bool {
	return true
}

func (n LeafNode) Weight() int {
	return n.RepNumber
}

func (n InternalNode) IsLeaf() bool {
	return true
}

func (n InternalNode) Weight() int {
	return n.RepNumber
}

func (n InternalNode) LeftChild() BaseNode {
	return n.LeftNode
}

func (n InternalNode) RightChild() BaseNode {
	return n.RightNode
}
