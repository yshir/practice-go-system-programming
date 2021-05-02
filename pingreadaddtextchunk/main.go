package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)

	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)

	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	// calclate CRC and add it
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())

	return &buffer
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

	newFile, err := os.Create("desktop_new.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunk(file)
	// write signature
	io.WriteString(newFile, "\x89PNG\r\nx1a\n")
	// write IHDR chunk to head
	io.Copy(newFile, chunks[0])
	// write text chunk
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	// write rest chunk
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	chunks2 := readChunk(newFile)
	for _, chunk := range chunks2 {
		dumpChunk(chunk)
	}
}
