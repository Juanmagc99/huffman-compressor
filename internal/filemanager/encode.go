package filemanager

import (
	"fmt"
	"os"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func WriteInFile(codeMap map[rune]string, path string) {
	var s string

	for k, v := range codeMap {
		s += fmt.Sprintf("%d:%s;", k, v)
	}
	s = s[0:len(s)-1] + "\n"

	fmt.Println(s)

	f, err := os.Create(path)
	utils.CheckError("Something went wrong with the creation of the file", err)

	n, err := f.Write([]byte(s))
	utils.CheckError("Somthing went wrong writing the file", err)
	fmt.Println(n)
}

func TextToCode(text string) string {
	return ""
}
