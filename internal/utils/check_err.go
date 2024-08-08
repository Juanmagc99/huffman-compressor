package utils

import (
	"fmt"
	"os"
)

/*
Given an error an a message it prints its output
if err != nil and stops the program
*/
func CheckError(message string, err error) {
	if err != nil {
		fmt.Println(message, ":", err)
		os.Exit(1)
	}
}
