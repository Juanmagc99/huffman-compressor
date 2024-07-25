package main

import (
	"fmt"
	"os"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func main() {

	filePath := "data/lorem_daa.txt"
	_, err := os.Open(filePath)

	utils.CheckError("Something", err)

	fmt.Println("Hola mundo!")

}
