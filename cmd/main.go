package main

import (
	"fmt"
	"os"
	"path"

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

	utils.TextToMap(f)

}
