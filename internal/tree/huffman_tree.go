package tree

import "fmt"

func CreateHuffmanTree(nl *NodeList) {
	for {
		if node, notEmpty := nl.PopFirst(); notEmpty {
			fmt.Printf("Character %d is poped from the array\n", node.(LeafNode).Character)
			fmt.Println(len(*nl))
		} else {
			break
		}
	}
}
