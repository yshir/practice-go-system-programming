package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunk(file *os.File) []io.Reader {
	var chunks []io.Reader

	// 最初の8バイト（signature）を skip
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err != nil {
			break
		}

		chunks = append(
			chunks,
			io.NewSectionReader(file, offset, int64(length)+12),
		)
		offset, _ = file.Seek(int64(length+8), 1)
	}

	return chunks
}

func main() {
	file, err := os.Open("desktop.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := readChunk(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
