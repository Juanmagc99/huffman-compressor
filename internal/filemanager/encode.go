package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func WriteInFile(codeMap *map[rune]string, f *os.File) {

	var s string
	var et string

	for k, v := range *codeMap {
		s += fmt.Sprintf("%d:%s;", k, v)
	}
	s = s[0:len(s)-1] + "\n"

	fmt.Println(strings.Split(f.Name(), ".")[0] + ".hoff")

	fNew, err := os.Create(strings.Split(f.Name(), ".")[0] + ".hoff")
	utils.CheckError("Something went wrong during file creation", err)

	WriteBits(s, true, fNew)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		cr := []rune(scanner.Text())[0]
		if len(et) >= 8 {
			WriteBits(et, false, fNew)
		} else {
			et += (*codeMap)[cr]
		}
	}

	if err := scanner.Err(); err != nil {
		utils.CheckError("Some problems scanning the file", err)
	}
}

func WriteBits(text string, isHeader bool, f *os.File) {
	f.WriteString(text)
}
