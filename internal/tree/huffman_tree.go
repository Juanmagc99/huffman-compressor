package tree

import (
	"sort"
)

/*
From a pointer to nodelist create a Huffman Tree
*/
func CreateHuffmanTree(nl *NodeList) {
	for {
		if len(*nl) > 1 {
			node1, _ := nl.PopFirst()
			node2, _ := nl.PopFirst()
			newNode := MergeNodes(node1, node2)

			sort.Slice(*nl, func(i, j int) bool {
				return (*nl)[i].Weight() < (*nl)[j].Weight()
			})

			*nl = append(*nl, newNode)
		} else {
			break
		}
	}
}

/*
Merge to structs that implements BaseNode into an InternalNode
*/
func MergeNodes(n1, n2 BaseNode) InternalNode {
	newNode := InternalNode{
		RepNumber: n1.Weight() + n2.Weight(),
		LeftNode:  n1,
		RightNode: n2,
	}

	return newNode
}

/*
Creates an Index of codes for every rune using the HuffmanTree(BaseNode)
*/
func IndexFromTree(tree BaseNode, code string, codeMap *map[rune]string) {
	if tree.IsLeaf() {
		(*codeMap)[tree.(LeafNode).Character] = code
	} else {
		IndexFromTree(tree.(InternalNode).LeftNode, code+"0", codeMap)
		IndexFromTree(tree.(InternalNode).RightNode, code+"1", codeMap)
	}
}
