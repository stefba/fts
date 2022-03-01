package main

import (
	"flag"
	"fmt"
)

type file struct {
	Path string
	Byte []byte
}

func main() {
	root := flag.String("path", ".", "path of text files to read")
	files, err := readFiles(*root)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(files)
}


