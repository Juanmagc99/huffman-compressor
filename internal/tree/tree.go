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

type NodeList []BaseNode

/*
Method for that implement a stack pop feature to array of nodes
*/
func (nl *NodeList) PopFirst() (BaseNode, bool) {
	if len(*nl) == 0 {
		return nil, false
	} else {
		index := len(*nl)
		node := (*nl)[0]
		*nl = (*nl)[1:index]
		return node, true
	}
}

func (n LeafNode) IsLeaf() bool {
	return true
}

func (n LeafNode) Weight() int {
	return n.RepNumber
}

func (n InternalNode) IsLeaf() bool {
	return false
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
