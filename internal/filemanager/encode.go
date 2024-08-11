package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

/*
Receive a file with it's codemap and create a file .hoff with the same name compressed
*/
func EncodeFile(codeMap *map[rune]string, f *os.File) {

	var s string
	et := ""

	for k, v := range *codeMap {
		s += fmt.Sprintf("%d:%s;", k, v)
	}
	s = s[0:len(s)-1] + "\n"

	fNew, err := os.Create(strings.Split(f.Name(), ".")[0] + ".huff")
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
		bValue := byte(0)
		for i, b := range text {
			if b == '1' {
				bValue |= (1 << (7 - i))
			}
		}
		_, err := f.Write([]byte{bValue})
		utils.CheckError("Error writing encode content", err)
	}

}
