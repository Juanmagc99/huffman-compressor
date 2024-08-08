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

	fNew, err := os.Create(strings.Split(f.Name(), ".")[0] + ".hoff")
	utils.CheckError("Something went wrong during file creation", err)

	defer fNew.Close()

	WriteBits(s, true, fNew)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		cr := []rune(scanner.Text())[0]
		et += (*codeMap)[cr]

		if len(et) >= 8 {
			WriteBits(et[:8], false, fNew)
			et = et[8:]
		}
	}

	if len(et) > 0 {
		et += strings.Repeat("0", 8-len(et))
		WriteBits(et, false, fNew)
	}

	if err := scanner.Err(); err != nil {
		utils.CheckError("Some problems scanning the file", err)
	}
}

/*
Receive a string with text and write it as the header directly if it is the map of codes
or transform strings of codes(len = 8) and push it as binary into a byte
*/
func WriteBits(text string, isHeader bool, f *os.File) {
	if isHeader {
		_, err := f.Write([]byte(text))
		utils.CheckError("Error writing header", err)
	} else {
		bValue := byte(0b0000000)
		for i, b := range text {
			if b == '1' {
				bValue |= (1 << (7 - i)) // OR logic with index of shift inverted (Little endian is set by default)
			}
		}
		_, err := f.Write([]byte{bValue})
		utils.CheckError("Error writing encode content", err)
	}

}
