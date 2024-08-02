package utils

import (
	"bufio"
	"os"
	"sort"
)

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

/*
This function receive a txt file and return a list of pairs with its runes
and the number of time it appears
*/
func TextToPairs(file *os.File) PairList {

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

	pl := make(PairList, len(cm))

	for k, v := range cm {
		pl = append(pl, Pair{k, v})
	}

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Value < pl[j].Value
	})

	return pl
}
