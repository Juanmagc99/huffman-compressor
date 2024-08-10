package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func DecodeFile(f *os.File) {
	/*

		Read first line, create map, read every byte trying to match codes in maps

	*/

	fNew, err := os.Create(strings.Split(f.Name(), ".")[0] + "_new.text")
	utils.CheckError("Something went wrong during file creation", err)

	defer fNew.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	headerString := ""
	//Read only header
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		headerString += scanner.Text()
	}
	fmt.Println(headerString)
}
