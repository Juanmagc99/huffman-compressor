package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com.Juanmagc99.huffman-compressor/internal/filemanager"
	"github.com.Juanmagc99.huffman-compressor/internal/tree"
	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func main() {

	compressFlag := flag.String("c", "", "File path to compress")
	decompressFlag := flag.String("d", "", "File path to decompress")

	flag.Parse()

	if *compressFlag != "" {
		filePath := *compressFlag

		if ext := path.Ext(filePath); ext != ".txt" {
			err := fmt.Errorf("program expects a .txt file extension not %s", ext)
			utils.CheckError("Problems with file extension", err)
		}

		f, err := os.Open(filePath)

		utils.CheckError("Something went wrong during file open", err)
		defer f.Close()

		nodeList := utils.TextToLeafs(f)

		tree.CreateHuffmanTree(&nodeList)

		huffmanTree, _ := nodeList.PopFirst()
		cm := make(map[rune]string)
		tree.IndexFromTree(huffmanTree, "", &cm)

		filemanager.EncodeFile(&cm, f)
		fmt.Println("File compressed successfully")
	}

	if *decompressFlag != "" {
		filePath := *decompressFlag

		if ext := path.Ext(filePath); ext != ".huff" {
			err := fmt.Errorf("program expects a .hoff file extension not %s", ext)
			utils.CheckError("Problems with file extension", err)
		}

		f, err := os.Open(filePath)

		utils.CheckError("Something went wrong during file open", err)
		defer f.Close()

		filemanager.DecodeFile(f)
		fmt.Println("File decompressed successfully")
	}

	if *compressFlag == "" && *decompressFlag == "" {
		fmt.Println("Please specify either -c (compress) or -d (decompress) with a file path.")
		flag.Usage()
	}

}
