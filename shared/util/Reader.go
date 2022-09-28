package util

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadAllByte(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
