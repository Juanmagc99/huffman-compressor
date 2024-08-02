package main

import (
	"fmt"
	"os"
	"path"

	"github.com.Juanmagc99.huffman-compressor/internal/tree"
	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func main() {

	filePath := "internal/testdata/lorem_data.txt"

	if ext := path.Ext(filePath); ext != ".txt" {
		err := fmt.Errorf("program expects a .txt file extension not %s", ext)
		utils.CheckError("Problems with file extension", err)
	}

	f, err := os.Open(filePath)

	utils.CheckError("Something went wrong during file open", err)
	defer f.Close()

	//pl := utils.TextToPairs(f)

	/*for _, p := range pl {
		fmt.Printf("The character %c appears %d times\n", p.Key, p.Value)
	}*/

	leafA := tree.LeafNode{'a', 5}
	leafB := tree.LeafNode{'b', 7}
	inter := tree.InternalNode{leafA, leafB, 12}
	inter2 := tree.InternalNode{leafA, inter, 17}

	fmt.Println(inter2)

}
