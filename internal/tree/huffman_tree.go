package tree

import "fmt"

func CreateHuffmanTree(leafList *LeafList) {
	for {
		if node, notEmpty := leafList.PopFirst(); notEmpty {
			fmt.Printf("Character %d is poped from the array\n", node.(LeafNode).Character)
			fmt.Println(len(*leafList))
		} else {
			break
		}
	}
}
