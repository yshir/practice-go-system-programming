package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file2, err2 := os.Create("b.txt")
	if err2 != nil {
		panic(err2)
	}
	io.Copy(file2, file)
}
