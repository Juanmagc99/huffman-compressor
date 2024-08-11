package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com.Juanmagc99.huffman-compressor/internal/utils"
)

func DecodeFile(f *os.File) {

	fNew, err := os.Create(strings.Split(f.Name(), ".")[0] + "_new.text")
	utils.CheckError("Something went wrong during file creation", err)

	defer fNew.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanBytes)

	cm := MapFromHeader(s)

	bStrings := ""

	for s.Scan() {
		b := s.Bytes()[0]
		bStrings += strconv.FormatUint(uint64(b), 2)
		WriteOnFile(&bStrings, cm, fNew)
	}

}

func WriteOnFile(s *string, cm map[string]rune, f *os.File) {

	length := len(*s)

	for i := 1; i <= length; i++ {
		if v, exists := cm[(*s)[:i]]; exists {
			fmt.Println(string(v))
			f.WriteString(string(v))
			*s = (*s)[i:length]
			length = len(*s)
		}
	}

}

func MapFromHeader(s *bufio.Scanner) map[string]rune {
	headerString := ""
	//Read only header
	for s.Scan() {
		if s.Text() == "\n" {
			break
		}
		headerString += s.Text()
	}

	invcm := make(map[string]rune)

	for _, rc := range strings.Split(headerString, ";") {
		rcSplit := strings.Split(rc, ":")
		runeVal, _ := strconv.Atoi(rcSplit[0])
		invcm[rcSplit[1]] = rune(runeVal)
	}

	return invcm
}
