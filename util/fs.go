package util

import (
	"bufio"
	"fmt"
	"os"
)

var cachedFiles map[string]*os.File

func ReadFile(path string) (*bufio.Scanner, error) {

	_, ok := cachedFiles[path]

	var file *os.File
	var err error

	if ok {
		file = cachedFiles[path]
	} else {
		file, err = os.Open(path)
	}

	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	file.Seek(0, 0)

	return bufio.NewScanner(file), nil
}
