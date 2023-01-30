package lib

import (
	"io/ioutil"
	"os"
)

func Load_file(filename string) []byte {
	file, openFileErr := os.Open(filename)
	if openFileErr != nil {
		panic(openFileErr)
	}
	defer file.Close()
	FileBin, readFileErr := ioutil.ReadAll(file)
	if readFileErr != nil {
		panic(readFileErr)
	}
	return FileBin
}
