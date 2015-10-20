package util

import (
	"io/ioutil"
	"os"
)

func ConfigBuffer(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return buf
}
