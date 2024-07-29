package utils

import (
	"bufio"
	"os"
)

/*
This function receive a txt file and return a map with its runes
and the number of time it appears
*/
func TextToMap(file *os.File) map[rune]int {

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

	return cm
}
