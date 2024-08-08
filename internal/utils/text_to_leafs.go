package utils

import (
	"bufio"
	"os"
	"sort"

	"github.com.Juanmagc99.huffman-compressor/internal/tree"
)

/*
This function receive a txt file and return a list of leaf with its runes
and their frecuency
*/
func TextToLeafs(f *os.File) tree.NodeList {

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	cm := make(map[rune]int)
	for scanner.Scan() {
		char := []rune(scanner.Text())[0]
		cm[char]++
	}

	if err := scanner.Err(); err != nil {
		CheckError("Some problems scanning the file", err)
	}

	//Point to the start of the file so its possible to scan it again
	_, err := f.Seek(0, 0)
	CheckError("Problem pointing to start of file", err)

	var nl tree.NodeList

	for k, v := range cm {
		nl = append(nl, tree.LeafNode{Character: k, RepNumber: v})
	}

	sort.Slice(nl, func(i, j int) bool {
		return nl[i].(tree.LeafNode).RepNumber < nl[j].(tree.LeafNode).RepNumber
	})

	return nl
}
