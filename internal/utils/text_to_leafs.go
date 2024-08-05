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
func TextToLeafs(file *os.File) tree.LeafList {

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	cm := make(map[rune]int)
	for scanner.Scan() {
		char := []rune(scanner.Text())[0]
		cm[char]++
	}

	if err := scanner.Err(); err != nil {
		CheckError("Some problems scanning the file", err)
	}

	var ll tree.LeafList

	for k, v := range cm {
		ll = append(ll, tree.LeafNode{Character: k, RepNumber: v})
	}

	sort.Slice(ll, func(i, j int) bool {
		return ll[i].RepNumber < ll[j].RepNumber
	})

	return ll
}
