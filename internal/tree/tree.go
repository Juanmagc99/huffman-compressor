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

type LeafList []LeafNode

func (ll *LeafList) PopFirst() (BaseNode, bool) {
	if len(*ll) == 0 {
		return nil, false
	} else {
		index := len(*ll)
		node := (*ll)[0]
		*ll = (*ll)[1:index]
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
