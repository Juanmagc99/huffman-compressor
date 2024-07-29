package main

import (
	"os"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func main() {

	filePath := "internal/testdata/lorem_data.txt"
	f, err := os.Open(filePath)

	utils.CheckError("Something", err)
	defer f.Close()

	utils.TextToMap(f)

}
