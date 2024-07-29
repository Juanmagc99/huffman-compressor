package utils

import (
	"fmt"
	"os"
)

func CheckError(message string, err error) {
	if err != nil {
		fmt.Println(message, ":", err)
		os.Exit(1)
	}
}
